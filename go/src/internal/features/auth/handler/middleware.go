package handler

import (
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/api/res"
	"net/http"
)

func CreateAuthMiddleware(service service.AuthService) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			accessToken, err := getAccessToken(r)

			if err != nil {
				res.WriteError(w, http.StatusUnauthorized, err, "no access token cookie")
				return
			}

			if err := service.VerifyToken(accessToken); err != nil {
				res.WriteError(w, http.StatusUnauthorized, err, "invalid access token")
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
