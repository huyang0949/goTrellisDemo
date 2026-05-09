package main

import (
	"context"
	"log"

	microinfra "goTrellisDemo/internal/infra/micro"
	"goTrellisDemo/internal/infra/resources"
	httptransport "goTrellisDemo/internal/transport/http"
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

	service := microinfra.NewService("api-gateway")
	greeterHandler := httptransport.NewGreeterHandler(service.Client(), "greeter")
	userHandler := httptransport.NewUserHandler(service.Client(), "user")
	router := httptransport.NewRouter(greeterHandler, userHandler)

	log.Println("api gateway listening on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
