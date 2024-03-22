package common

import (
	"mmt/mmt/model"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

var (
	ErrorNone         error = errors.New("token is none! ")
	ErrorExpired      error = errors.New("token is expired! ")
	ErrorMalformed    error = errors.New("token format error! ")
	ErrorNotValidYet  error = errors.New("token not valid yet! ")
	ErrorInvalid      error = errors.New("token invalid! ")
	ErrInvalidKeyType error = errors.New("key is of invalid type! ")
)

func CreateJwtToken(secretKey string, exp, now int64, user model.MmtUsers) (string, error) {
	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp + now,
			Issuer:    "mmt",
		},
		RefreshAfter: (now + exp) - 3600, // 提前1h
		UserInfo:     user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
func ParseToken(secretKey, AuthToken string) (*model.Claims, error) {
	if token, err := jwt.ParseWithClaims(AuthToken, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}); err != nil || token == nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrorExpired
			} else if e.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrorMalformed
			} else if e.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrorNotValidYet
			} else {
				return nil, ErrInvalidKeyType
			}
		default:
			return nil, ErrorInvalid
		}
	} else {
		if v, ok := token.Claims.(*model.Claims); ok && token.Valid {
			return v, nil
		}
	}
	return nil, nil
}
