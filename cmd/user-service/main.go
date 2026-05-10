package main

import (
	"context"
	"log"

	appuser "goTrellisDemo/internal/app/user"
	microinfra "goTrellisDemo/internal/infra/micro"
	"goTrellisDemo/internal/infra/resources"
	grpcuser "goTrellisDemo/internal/transport/grpc/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	ctx := context.Background()
	res, err := resources.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := res.Close(ctx); err != nil {
			log.Printf("close resources: %v", err)
		}
	}()

	service := microinfra.NewService(
		"user",
		micro.Address(":8082"),
	)

	if err := grpcuser.Register(service.Server(), appuser.NewService(res.Mongo.Database)); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
