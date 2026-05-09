package greeter

import "context"

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello " + req.Name}, nil
}
