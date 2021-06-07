package main

import (
	"fmt"
	"github.com/adhimaswaskita/go-interface/handler"
	service "github.com/adhimaswaskita/go-interface/service"
)

func main() {
	service := service.Service{}
	handler := handler.NewHandler(service)

	fmt.Printf("\nHello from main : %v\n", handler.Hello())
}