package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const userCollectionName = "user"

type Service struct {
	users *mongo.Collection
}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	IsNew         bool   `json:"isNew"`
	RegisterTime  int64  `json:"registerTime"`
	LastLoginTime int64  `json:"lastLoginTime"`
}

type userRecord struct {
	UserId        string `bson:"userId"`
	Username      string `bson:"username"`
	Avatar        string `bson:"avatar"`
	RegisterTime  int64  `bson:"registerTime"`
	LastLoginTime int64  `bson:"lastLoginTime"`
}

func NewService(database *mongo.Database) *Service {
	return &Service{
		users: database.Collection(userCollectionName),
	}
}

func (s *Service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	if req.UserId == "" {
		return nil, fmt.Errorf("userId is required")
	}

	now := time.Now().Unix()
	filter := bson.M{"userId": req.UserId}

	var existing userRecord
	if err := s.users.FindOne(ctx, filter).Decode(&existing); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("find user: %w", err)
		}

		record := userRecord{
			UserId:        req.UserId,
			Username:      req.Username,
			Avatar:        req.Avatar,
			RegisterTime:  now,
			LastLoginTime: now,
		}
		if _, err := s.users.InsertOne(ctx, record); err != nil {
			return nil, fmt.Errorf("insert user: %w", err)
		}

		return &LoginResponse{
			Token:         "token",
			IsNew:         true,
			RegisterTime:  record.RegisterTime,
			LastLoginTime: record.LastLoginTime,
		}, nil
	}

	update := bson.M{
		"$set": bson.M{
			"lastLoginTime": now,
		},
	}
	setFields := update["$set"].(bson.M)
	if existing.Username != req.Username {
		setFields["username"] = req.Username
	}
	if existing.Avatar != req.Avatar {
		setFields["avatar"] = req.Avatar
	}

	if _, err := s.users.UpdateOne(ctx, filter, update); err != nil {
		return nil, fmt.Errorf("update user: %w", err)
	}

	return &LoginResponse{
		Token:         "token",
		IsNew:         false,
		RegisterTime:  existing.RegisterTime,
		LastLoginTime: now,
	}, nil
}
