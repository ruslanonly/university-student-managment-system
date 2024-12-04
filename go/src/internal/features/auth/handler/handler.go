package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	"github.com/ruslanonly/university-student-managment-system/src/pkg/api/res"
	"log/slog"
	"net/http"
)

type Handler struct {
	log         *slog.Logger
	authService *service.AuthService
	cfg         *auth.Config
}

// OAuth godoc
// @Summary Начать авторизацию по протоколу OAuth 2.0
// @Description Начать авторизацию по протоколу OAuth 2.0
// @Tags OAuth
// @Param provider path string true "Название провайдера"
// @Accept json
// @Produce json
// @Router /auth/{provider} [get]
// @Success 200
func (h *Handler) OAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)

	if err != nil {
		gothic.BeginAuthHandler(w, r)
		return
	}

	res.WriteJSON(w, http.StatusOK, user)
}

// OAuthCallback godoc
// @Summary Завершить авторизацию по протоколу OAuth 2.0
// @Description Завершить авторизацию по протоколу OAuth 2.0
// @Tags OAuth
// @Param provider path string true "Название провайдера"
// @Accept json
// @Produce json
// @Router /auth/{provider}/callback [get]
// @Success 200
func (h *Handler) OAuthCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)

	if err != nil {
		res.WriteError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	h.log.Info("user", user)

	accessToken, err := h.authService.Login(r.Context(), service.LoginDTO{Username: user.Email})

	if err != nil {
		res.WriteError(w, http.StatusInternalServerError, err, err.Error())
		return
	}

	setAccessToken(w, h.cfg, accessToken)

	http.Redirect(w, r, h.cfg.RedirectURL, http.StatusPermanentRedirect)
}

func New(log *slog.Logger, cfg *auth.Config, authService *service.AuthService) *Handler {
	return &Handler{
		log:         log,
		authService: authService,
		cfg:         cfg,
	}
}

func Init(r *chi.Mux, h *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Get("/{provider}", h.OAuth)
		r.Get("/{provider}/callback", h.OAuthCallback)
	})
}
