package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

type Greeter struct{}

type GreeterRequest struct {
	Name string `json:"name"`
}

type GreeterResponse struct {
	Message string `json:"message"`
}

func (g *Greeter) Hello(ctx context.Context, req *GreeterRequest, rsp *GreeterResponse) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Address(":8080"),
		micro.Version("latest"),
	)
	service.Init()

	micro.RegisterHandler(service.Server(), new(Greeter))

	go func() {
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	router.POST("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
			return
		}

		req := service.Client().NewRequest(
			"greeter",
			"Greeter.Hello",
			&GreeterRequest{Name: name},
			client.WithContentType("application/json"),
		)
		rsp := new(GreeterResponse)
		if err := service.Client().Call(c.Request.Context(), req, rsp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, rsp)
	})

	log.Println("api gateway listening on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
