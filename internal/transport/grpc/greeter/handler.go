package greeter

import (
	"context"

	appgreeter "goTrellisDemo/internal/app/greeter"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type Greeter struct {
	service *appgreeter.Service
}

func NewGreeter(service *appgreeter.Service) *Greeter {
	return &Greeter{service: service}
}

func Register(s server.Server, service *appgreeter.Service) error {
	return micro.RegisterHandler(s, NewGreeter(service))
}

func (g *Greeter) Hello(ctx context.Context, req *appgreeter.HelloRequest, rsp *appgreeter.HelloResponse) error {
	result, err := g.service.Hello(ctx, req)
	if err != nil {
		return err
	}

	*rsp = *result
	return nil
}
