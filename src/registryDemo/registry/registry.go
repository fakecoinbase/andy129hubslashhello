package registry

import "context"

// 服务注册发现的总接口
type Registry interface {
	// 1. 初始化问题
	Init(ctx context.Context, opts ...Option) (err error)
	// 2.Name
	Name() string
	// 3.注册
	Registry(ctx context.Context, service *Service) (err error)
	// 4.反注册
	UnRegistry(ctx context.Context, service *Service) (err error)
	// 5.服务发现
	GetService(ctx context.Context, name string) (service *Service, err error)
}
