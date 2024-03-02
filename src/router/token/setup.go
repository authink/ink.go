package token

import (
	"database/sql"
	errs "errors"
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone/web"
	"github.com/authink/inkstone/util"
	"github.com/gin-gonic/gin"
)

func SetupTokenGroup(rg *gin.RouterGroup) {
	gToken := rg.Group("token")
	gToken.POST("grant", web.HandlerAdapter(grant))
	gToken.POST("refresh", web.HandlerAdapter(refresh))
	gToken.POST("revoke", web.HandlerAdapter(revoke))
}

func generateAuthToken(c *web.Context, app *model.App, staff *model.Staff) (res *GrantRes) {
	appCtx := c.AppContext()

	jwtClaims := util.NewJwtClaims(
		util.GenerateUUID(),
		appCtx.AppName,
		app.Name,
		staff.Email,
		time.Duration(appCtx.AccessTokenDuration),
		app.Id,
		staff.Id,
	)

	accessToken, err := util.GenerateToken(
		appCtx.SecretKey,
		jwtClaims,
	)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	refreshToken := util.GenerateUUID()

	authToken := model.NewAuthToken(jwtClaims.ID, refreshToken, app.Id, staff.Id)

	if err = orm.AuthToken(appCtx).Insert(authToken); err != nil {
		c.AbortWithServerError(err)
		return
	}

	res = &GrantRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int(appCtx.AccessTokenDuration),
	}
	return
}

func checkRefreshToken(c *web.Context, refreshToken string) (authToken *model.AuthToken, ok bool) {
	appCtx := c.AppContext()
	authToken, err := orm.AuthToken(appCtx).GetByRefreshToken(refreshToken)
	if err != nil {
		if !errs.Is(err, sql.ErrNoRows) {
			c.AbortWithServerError(err)
			return
		}
		c.AbortWithClientError(errors.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	if err = orm.AuthToken(appCtx).Delete(int(authToken.Id)); err != nil {
		c.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(time.Duration(appCtx.RefreshTokenDuration) * time.Hour)) {
		c.AbortWithClientError(errors.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	ok = true
	return
}
