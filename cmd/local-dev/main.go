package main

import (
	"log"

	appgreeter "goTrellisDemo/internal/app/greeter"
	appuser "goTrellisDemo/internal/app/user"
	microinfra "goTrellisDemo/internal/infra/micro"
	grpcgreeter "goTrellisDemo/internal/transport/grpc/greeter"
	grpcuser "goTrellisDemo/internal/transport/grpc/user"
	httptransport "goTrellisDemo/internal/transport/http"

	"github.com/micro/go-micro/v2"
)

func main() {
	backendService := microinfra.NewService(
		"local-backend",
		micro.Address(":8080"),
	)
	if err := grpcgreeter.Register(backendService.Server(), appgreeter.NewService()); err != nil {
		log.Fatal(err)
	}
	if err := grpcuser.Register(backendService.Server(), appuser.NewService()); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := backendService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	gatewayService := microinfra.NewService("api-gateway")
	greeterHandler := httptransport.NewGreeterHandler(gatewayService.Client(), "local-backend", "127.0.0.1:8080")
	userHandler := httptransport.NewUserHandler(gatewayService.Client(), "local-backend", "127.0.0.1:8080")
	router := httptransport.NewRouter(greeterHandler, userHandler)

	log.Println("local dev api gateway listening on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
