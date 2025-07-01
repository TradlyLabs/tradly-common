package service

import "context"

type Service interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type ServiceManager interface {
	Register(name string, svc Service)
}
