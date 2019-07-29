package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var jwtSecret []byte

func GenerateTocken(username, password string) (string, error) {
	nowTime := time.Now()

	expireTime := nowTime.Add(24 * 30 * time.Hour)

	Claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tqgin",
		},
	}

	tokenClainms := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tocken, err := tokenClainms.SignedString(jwtSecret)
	return tocken, err
}

func ParseToken(tocken string) (*Claims, error) {
	tokenClainms, err := jwt.ParseWithClaims(tocken, &Claims{}, func(tocken *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClainms != nil {
		if Claims, ok := tokenClainms.Claims.(*Claims); ok && tokenClainms.Valid {
			return Claims, nil
		}
	}
	return nil, err
}
