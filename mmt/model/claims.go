package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	RefreshAfter int64    `json:"refresh_after"`
	UserInfo     MmtUsers `json:"user_info"`
}
