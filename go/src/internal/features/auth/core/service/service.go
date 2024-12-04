package service

import (
	"context"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
)

type AuthService struct {
	cfg *auth.Config
}

func New(
	cfg *auth.Config,

) *AuthService {
	return &AuthService{
		cfg: cfg,
	}
}

func (s *AuthService) Login(_ context.Context, dto LoginDTO) (model.AccessToken, error) {
	// Создание AccessToken
	accessToken, err := model.CreateAccessToken(s.cfg.Secret, dto.Username, s.cfg.AccessTokenTTL)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *AuthService) VerifyToken(token model.AccessToken) error {
	return model.VerifyAccessToken(token, s.cfg.Secret)
}
