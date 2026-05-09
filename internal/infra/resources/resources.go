package resources

import (
	"context"
	"fmt"
	"log"

	"goTrellisDemo/internal/config"
	mongoinfra "goTrellisDemo/internal/infra/mongo"
	redisinfra "goTrellisDemo/internal/infra/redis"
)

type Resources struct {
	Config *config.Config
	Mongo  *mongoinfra.Client
	Redis  *redisinfra.Client
}

func New(ctx context.Context) (*Resources, error) {
	cfg, err := config.Load("")
	if err != nil {
		return nil, err
	}

	mongoClient, err := mongoinfra.NewClient(ctx, cfg.MongoDBURL)
	if err != nil {
		return nil, err
	}

	redisClient, err := redisinfra.NewClient(cfg.Redis)
	if err != nil {
		_ = mongoClient.Close(ctx)
		return nil, err
	}

	log.Printf("mongodb configured: %s", cfg.MongoDBURL)
	log.Printf("redis configured: %s ttl=%s", cfg.Redis.Addr(), cfg.Redis.TTL())

	return &Resources{
		Config: cfg,
		Mongo:  mongoClient,
		Redis:  redisClient,
	}, nil
}

func (r *Resources) Close(ctx context.Context) error {
	if r == nil {
		return nil
	}

	var closeErr error
	if err := r.Redis.Close(); err != nil {
		closeErr = fmt.Errorf("close redis: %w", err)
	}
	if err := r.Mongo.Close(ctx); err != nil {
		if closeErr != nil {
			closeErr = fmt.Errorf("%v; close mongo: %w", closeErr, err)
		} else {
			closeErr = fmt.Errorf("close mongo: %w", err)
		}
	}

	return closeErr
}
