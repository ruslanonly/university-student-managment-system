package service

import (
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/repository"
)

type AuthService struct {
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func New(
	userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}
