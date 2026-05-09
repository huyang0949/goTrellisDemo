package main

import (
	"log"

	appgreeter "goTrellisDemo/internal/app/greeter"
	microinfra "goTrellisDemo/internal/infra/micro"
	grpcgreeter "goTrellisDemo/internal/transport/grpc/greeter"
	httptransport "goTrellisDemo/internal/transport/http"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/memory"
)

func main() {
	registry := memory.NewRegistry()
	greeterService := microinfra.NewService(
		"greeter",
		micro.Address(":8080"),
		micro.Registry(registry),
	)
	if err := grpcgreeter.Register(greeterService.Server(), appgreeter.NewService()); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := greeterService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	gatewayService := microinfra.NewService(
		"api-gateway",
		micro.Registry(registry),
	)
	greeterHandler := httptransport.NewGreeterHandler(gatewayService.Client(), "greeter", "127.0.0.1:8080")
	router := httptransport.NewRouter(greeterHandler)

	log.Println("local dev api gateway listening on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
