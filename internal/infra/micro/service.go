package microinfra

import "github.com/micro/go-micro/v2"

func NewService(name string, opts ...micro.Option) micro.Service {
	options := []micro.Option{
		micro.Name(name),
		micro.Version("latest"),
	}
	options = append(options, opts...)

	service := micro.NewService(options...)
	service.Init()
	return service
}
