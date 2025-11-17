package handlers

import "hw3/internal/application/service"

type handler struct {
	service service.AccountService
}

func NewHandler(service service.AccountService) handler {
	return handler{service: service}
}
