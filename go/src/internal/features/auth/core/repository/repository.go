package repository

import (
	"context"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
)

type GetSessionDTO struct {
	UserID int `json:"user_id"`
}

type SessionRepository interface {
	Get(ctx context.Context, dto GetSessionDTO) (*model.Session, error)
	Save(ctx context.Context, session *model.Session) (*model.Session, error)
}
