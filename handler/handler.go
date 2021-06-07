package handler

import "github.com/adhimaswaskita/go-interface/service"

type IHandler interface {
	Hello() string
}

type Handler struct {
	Service service.IService
}

func NewHandler(s service.IService) IHandler {
	handler := &Handler{
		Service: s,
	}

	return handler
}