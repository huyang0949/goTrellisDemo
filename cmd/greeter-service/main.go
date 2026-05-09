package main

import (
	"log"

	appgreeter "goTrellisDemo/internal/app/greeter"
	microinfra "goTrellisDemo/internal/infra/micro"
	grpcgreeter "goTrellisDemo/internal/transport/grpc/greeter"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := microinfra.NewService(
		"greeter",
		micro.Address(":8080"),
	)

	if err := grpcgreeter.Register(service.Server(), appgreeter.NewService()); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
