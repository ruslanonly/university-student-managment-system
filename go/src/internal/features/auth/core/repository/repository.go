package repository

import (
	"context"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
)

type GetUserDTO struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
}

type UserRepository interface {
	Get(ctx context.Context, dto GetUserDTO) (*model.User, error)
	Save(ctx context.Context, user *model.User) (*model.User, error)
}

type GetSessionDTO struct {
	UserID int `json:"user_id"`
}

type SessionRepository interface {
	Get(ctx context.Context, dto GetSessionDTO) (*model.Session, error)
	Save(ctx context.Context, session *model.Session) (*model.Session, error)
}
