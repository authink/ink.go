package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

type authToken interface {
	ORM[model.AuthToken]
	GetByRefreshToken(string) (*model.AuthToken, error)
	GetByAccessToken(string) (*model.AuthToken, error)
}

type authTokenImpl core.Ink

// Find implements authToken.
func (*authTokenImpl) Find() ([]model.AuthToken, error) {
	panic("unimplemented")
}

// Get implements authToken.
func (*authTokenImpl) Get(int) (*model.AuthToken, error) {
	panic("unimplemented")
}

// SaveWithTx implements authToken.
func (*authTokenImpl) SaveWithTx(*model.AuthToken, *sqlx.Tx) error {
	panic("unimplemented")
}

// GetByRefreshToken implements authToken.
func (at *authTokenImpl) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = at.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// GetByAccessToken implements authToken.
func (at *authTokenImpl) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = at.DB.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// Save implements authToken.
func (at *authTokenImpl) Save(token *model.AuthToken) (err error) {
	_, err = at.DB.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
	return
}

// Delete implements authToken.
func (at *authTokenImpl) Delete(id int) (err error) {
	_, err = at.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(ink *core.Ink) authToken {
	return (*authTokenImpl)(ink)
}
