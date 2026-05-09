package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Client struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewClient(ctx context.Context, uri string) (*Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("connect mongo: %w", err)
	}

	dbName, err := databaseName(uri)
	if err != nil {
		_ = client.Disconnect(ctx)
		return nil, err
	}

	return &Client{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c == nil || c.Client == nil {
		return nil
	}
	return c.Client.Disconnect(ctx)
}
