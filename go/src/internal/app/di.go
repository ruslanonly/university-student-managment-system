package app

import authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"

type diContainer struct {
	authHandler *authHandler.Handler
}

func (a *App) inject() {
	container := &diContainer{
		authHandler: authHandler.New(a.log),
	}
	a.container = container
}
