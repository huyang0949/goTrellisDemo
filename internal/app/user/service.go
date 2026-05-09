package user

import (
	"context"
)

type Service struct{}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Avatar   string `json:"password"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	IsNew         bool   `json:"isNew"`
	RegisterTime  int64  `json:"registerTime"`
	LastLoginTime int64  `json:"lastLoginTime"`
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{
		Token:         "token",
		IsNew:         true,
		RegisterTime:  0,
		LastLoginTime: 0,
	}, nil
}
