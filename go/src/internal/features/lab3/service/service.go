package service

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service struct {
	elasticCli *elasticsearch.Client
	neoCli     neo4j.DriverWithContext
	pgCli      *pgx.Conn
	redisCli   *redis.Client
	mongoCli   *mongo.Client
}

func (s *Service) Execute(ctx context.Context, in *In) (*Out, error) {
	// TODO implement for lab 3
	return &Out{}, nil
}

func New(elasticCli *elasticsearch.Client, neoCli neo4j.DriverWithContext, pgCli *pgx.Conn, redisCli *redis.Client, mongoCli *mongo.Client) *Service {
	return &Service{
		elasticCli: elasticCli,
		neoCli:     neoCli,
		pgCli:      pgCli,
		redisCli:   redisCli,
		mongoCli:   mongoCli,
	}
}
