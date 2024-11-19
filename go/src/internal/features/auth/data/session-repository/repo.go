package session_repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/model"
	"github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/repository"
)

type SessionRepository struct {
	conn *pgx.Conn
}

func (s SessionRepository) Get(ctx context.Context, dto repository.GetSessionDTO) (*model.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionRepository) Save(ctx context.Context, session *model.Session) (*model.Session, error) {
	query := `
		INSERT INTO sessions (username, token, expires_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (username) DO UPDATE
		SET username = $1, token = $2, expires_at = $3
	`

	_, err := s.conn.Exec(ctx, query, session.Username, session.Token, session.ExpiresAt)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func New(conn *pgx.Conn) repository.SessionRepository {
	return SessionRepository{
		conn: conn,
	}
}
