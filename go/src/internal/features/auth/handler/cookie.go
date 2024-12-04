package handler

import (
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
	"net/http"
)

const (
	accessTokenCookieName = "access_token"
)

func getAccessToken(r *http.Request) (model.AccessToken, error) {
	cookie, err := r.Cookie(accessTokenCookieName)

	if err != nil {
		return "", err
	}

	return model.AccessToken(cookie.Value), nil
}

func setAccessToken(w http.ResponseWriter, cfg *auth.Config, accessToken model.AccessToken) {
	accessTokenCookie := http.Cookie{
		Name:     accessTokenCookieName,
		Value:    string(accessToken),
		Path:     "/",
		MaxAge:   cfg.AccessTokenTTL,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &accessTokenCookie)
}
