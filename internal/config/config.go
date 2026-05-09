package config

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const DefaultPath = "configs/local.json"

type Config struct {
	MongoDBURL string      `json:"mongodbUrl"`
	Redis      RedisConfig `json:"redis"`
}

type RedisConfig struct {
	Mode       string             `json:"mode"`
	Cluster    bool               `json:"cluster"`
	Port       int                `json:"port"`
	Host       string             `json:"host"`
	DefaultCfg RedisDefaultConfig `json:"defaultCfg"`
}

type RedisDefaultConfig struct {
	Redis RedisTTLConfig `json:"redis"`
}

type RedisTTLConfig struct {
	TTL int `json:"ttl"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = os.Getenv("APP_CONFIG")
	}
	if path == "" {
		path = DefaultPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %q: %w", path, err)
	}

	return &cfg, nil
}

func (c RedisConfig) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

func (c RedisConfig) TTL() time.Duration {
	return time.Duration(c.DefaultCfg.Redis.TTL) * time.Second
}
