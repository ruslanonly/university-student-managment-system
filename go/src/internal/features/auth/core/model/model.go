package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("[ErrInvalidToken]")
)

type AccessToken string

func CreateAccessToken(secretKey string, userID string, ttl int) (AccessToken, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "university-student-management-system",
		"exp": time.Now().Add(time.Duration(ttl)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return AccessToken(tokenString), nil
}

func VerifyAccessToken(accessToken AccessToken, secret string) error {
	token, err := jwt.Parse(string(accessToken), func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return errors.Wrap(ErrInvalidToken, err.Error())
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}
