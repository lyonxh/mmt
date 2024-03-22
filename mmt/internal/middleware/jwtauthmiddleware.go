package middleware

import (
	"context"
	"errors"
	"mmt/utils"
	"net/http"
	"time"
)

type Claims string
type JwtAuthMiddleware struct {
	AccessSecret string
}

func NewJwtAuthMiddleware(AccessSecret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		AccessSecret,
	}

}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			common.ResResult(w, nil, common.ErrorNone)
			return
		}
		token, err := common.ParseToken(m.AccessSecret, r.Header.Get("Authorization"))
		if err != nil {
			common.ResResult(w, nil, err)
			return
		}
		now := time.Now().Unix()
		if now >= token.RefreshAfter {
			if refreshToken, err := common.CreateJwtToken(m.AccessSecret, token.ExpiresAt, now, token.UserInfo); err != nil {
				common.ResResult(w, nil, errors.New("refresh token error! "))
				return
			} else {
				w.Header().Set("refresh-token", refreshToken)
			}
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, Claims("claims"), token.UserInfo)
		next(w, r.WithContext(ctx))
	}
}
