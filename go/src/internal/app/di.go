package app

import (
	"github.com/elastic/go-elasticsearch/v8"
	authService "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"
	lab1Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/handler"
	lab1Service "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/service"
)

type diContainer struct {
	authHandler *authHandler.Handler
	lab1Handler *lab1Handler.Handler
}

func (a *App) inject() func() {
	//conn, err := pgx.Connect(context.Background(), a.cfg.Postgres.ConnectionString)
	//
	//if err != nil {
	//	panic(err)
	//}

	elasticCli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{a.cfg.Elastic.ConnectionString},
	})

	if err != nil {
		panic(err)
	}

	authServiceImpl := authService.New(&a.cfg.Auth)

	lab1ServiceImpl := lab1Service.New(elasticCli)

	container := &diContainer{
		authHandler: authHandler.New(a.log, &a.cfg.Auth, authServiceImpl),
		lab1Handler: lab1Handler.New(lab1ServiceImpl),
	}

	a.container = container

	return func() {
		//_ = conn.Close(context.Background())
	}
}
