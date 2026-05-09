package user

import (
	"context"

	appuser "goTrellisDemo/internal/app/user"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type User struct {
	service *appuser.Service
}

func NewUser(service *appuser.Service) *User {
	return &User{service: service}
}

func Register(s server.Server, service *appuser.Service) error {
	return micro.RegisterHandler(s, NewUser(service))
}

func (u *User) Login(ctx context.Context, req *appuser.LoginRequest, rsp *appuser.LoginResponse) error {
	result, err := u.service.Login(ctx, req)
	if err != nil {
		return err
	}

	*rsp = *result
	return nil
}
