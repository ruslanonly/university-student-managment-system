package app

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	authService "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/core/service"
	authHandler "github.com/ruslanonly/university-student-managment-system/src/internal/features/auth/handler"
	lab1Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/handler"
	lab1Service "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab1/service"
	lab2Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab2/handler"
	lab2Service "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab2/service"
	lab3Handler "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab3/handler"
	lab3Service "github.com/ruslanonly/university-student-managment-system/src/internal/features/lab3/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type diContainer struct {
	authService *authService.AuthService
	authHandler *authHandler.Handler
	lab1Handler *lab1Handler.Handler
	lab2Handler *lab2Handler.Handler
	lab3Handler *lab3Handler.Handler
}

func (a *App) inject(ctx context.Context) func() {
	pgCli, err := pgx.Connect(context.Background(), a.cfg.Postgres.ConnectionString)

	if err != nil {
		panic(err)
	}

	elasticCli, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{a.cfg.Elastic.ConnectionString},
	})

	if err != nil {
		panic(err)
	}

	neoCli, err := neo4j.NewDriverWithContext(
		a.cfg.Neo.URI,
		neo4j.BasicAuth(a.cfg.Neo.User, a.cfg.Neo.Password, ""),
	)

	if err != nil {
		panic(err)
	}

	err = neoCli.VerifyConnectivity(ctx)

	if err != nil {
		panic(err)
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:     a.cfg.Redis.Host,
		Username: a.cfg.Redis.User,
		Password: a.cfg.Redis.Password,
		DB:       a.cfg.Redis.DB,
	})

	mongoServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoOpts := options.Client().ApplyURI(a.cfg.Mongo.Address).SetServerAPIOptions(mongoServerAPI)

	mongoCli, err := mongo.Connect(mongoOpts)

	if err != nil {
		panic(err)
	}

	err = mongoCli.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	authServiceImpl := authService.New(&a.cfg.Auth)

	lab1ServiceImpl := lab1Service.New(elasticCli, neoCli, pgCli, redisCli)

	lab2ServiceImpl := lab2Service.New(elasticCli, neoCli, pgCli, redisCli, mongoCli)

	lab3ServiceImpl := lab3Service.New(elasticCli, neoCli, pgCli, redisCli, mongoCli)

	container := &diContainer{
		authService: authServiceImpl,
		authHandler: authHandler.New(a.log, &a.cfg.Auth, authServiceImpl),
		lab1Handler: lab1Handler.New(lab1ServiceImpl),
		lab2Handler: lab2Handler.New(lab2ServiceImpl),
		lab3Handler: lab3Handler.New(lab3ServiceImpl),
	}

	a.container = container

	return func() {
		_ = pgCli.Close(context.Background())
		_ = redisCli.Close()
		_ = mongoCli.Disconnect(ctx)
	}
}
