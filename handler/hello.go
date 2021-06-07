package handler

import "fmt"

func(h *Handler) Hello() string {
	sValue := h.Service.Hello();
	fmt.Printf("Return from handler : %v", sValue)
	return sValue
}