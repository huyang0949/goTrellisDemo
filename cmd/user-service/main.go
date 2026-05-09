package main

import (
	"log"

	appuser "goTrellisDemo/internal/app/user"
	microinfra "goTrellisDemo/internal/infra/micro"
	grpcuser "goTrellisDemo/internal/transport/grpc/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := microinfra.NewService(
		"user",
		micro.Address(":8082"),
	)

	if err := grpcuser.Register(service.Server(), appuser.NewService()); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
