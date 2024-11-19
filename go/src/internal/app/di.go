package app

import (
	"context"
	"github.com/jackc/pgx/v4"
	authService "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	sessionRepository "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/data/session-repository"
	authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"
)

type diContainer struct {
	authHandler *authHandler.Handler
}

func (a *App) inject() func() {
	conn, err := pgx.Connect(context.Background(), a.cfg.Postgres.ConnectionString)

	if err != nil {
		panic(err)
	}

	sessionRepositoryImpl := sessionRepository.New(conn)

	authServiceImpl := authService.New(sessionRepositoryImpl, &a.cfg.Auth)

	container := &diContainer{
		authHandler: authHandler.New(a.log, &a.cfg.Auth, authServiceImpl),
	}

	a.container = container

	return func() {
		_ = conn.Close(context.Background())
	}
}
