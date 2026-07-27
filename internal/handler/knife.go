package handler

import "knives/internal/service"

type KnifeHandler struct {
	service service.KnifeServiceInterface
}

func NewKnifeHandler(service service.KnifeServiceInterface) *KnifeHandler {
	return &KnifeHandler{service: service}
}
