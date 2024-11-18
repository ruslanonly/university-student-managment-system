package model

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("[ErrInvalidToken]")
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RefreshToken string

type Session struct {
	UserID    int          `json:"user_id"`
	Token     RefreshToken `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
}

func NewSession(userID int, ttl int) *Session {
	return &Session{
		UserID:    userID,
		Token:     RefreshToken(uuid.New().String()),
		ExpiresAt: time.Now().Add(time.Duration(ttl)),
	}
}

type AccessToken string

func CreateAccessToken(secretKey string, username string, ttl int) (AccessToken, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "university-student-management-system",
		"exp": time.Now().Add(time.Duration(ttl)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Printf("Token claims added: %+v\n", claims)
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
