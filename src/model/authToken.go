package model

import "github.com/authink/inkstone"

type AuthToken struct {
	inkstone.Model
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	AppId        uint32 `db:"app_id"`
	AccountId    uint32 `db:"account_id"`
}

func NewAuthToken(accessToken, refreshToken string, appId, accountId uint32) *AuthToken {
	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AppId:        appId,
		AccountId:    accountId,
	}
}

type AuthTokenWithApp struct {
	inkstone.Model
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	AppId        uint32 `db:"app_id"`
	AppName      string `db:"app_name"`
	AccountId    uint32 `db:"account_id"`
}
