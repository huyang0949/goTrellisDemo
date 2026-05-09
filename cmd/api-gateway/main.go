package main

import (
	"log"

	microinfra "goTrellisDemo/internal/infra/micro"
	httptransport "goTrellisDemo/internal/transport/http"
)

func main() {
	service := microinfra.NewService("api-gateway")
	greeterHandler := httptransport.NewGreeterHandler(service.Client(), "greeter")
	router := httptransport.NewRouter(greeterHandler)

	log.Println("api gateway listening on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
