package redis

import (
	"context"
	"fmt"

	"goTrellisDemo/internal/config"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Client *redis.Client
	TTL    int
}

func NewClient(cfg config.RedisConfig) (*Client, error) {
	if cfg.Cluster {
		return nil, fmt.Errorf("redis cluster mode is not supported yet")
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr(),
	})

	return &Client{
		Client: client,
		TTL:    cfg.DefaultCfg.Redis.TTL,
	}, nil
}

func (c *Client) Close() error {
	if c == nil || c.Client == nil {
		return nil
	}
	return c.Client.Close()
}

func (c *Client) Ping(ctx context.Context) error {
	if c == nil || c.Client == nil {
		return nil
	}
	return c.Client.Ping(ctx).Err()
}
