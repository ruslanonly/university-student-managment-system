package service

import (
	"context"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/repository"
)

type AuthService struct {
	sessionRepository repository.SessionRepository
	cfg               *auth.Config
}

func New(
	sessionRepository repository.SessionRepository,
	cfg *auth.Config,

) *AuthService {
	return &AuthService{
		sessionRepository: sessionRepository,
		cfg:               cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, dto LoginDTO) (model.AccessToken, model.RefreshToken, error) {
	// Создание AccessToken
	accessToken, err := model.CreateAccessToken(s.cfg.Secret, dto.Username, s.cfg.AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	// Создание и сохранение сессии пользователя
	session := model.NewSession(dto.Username, s.cfg.RefreshTokenTTL)

	session, err = s.sessionRepository.Save(ctx, session)
	if err != nil {
		return "", "", err
	}

	return accessToken, session.Token, nil
}
