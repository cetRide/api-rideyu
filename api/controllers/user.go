package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *UseCaseHandler) Register(http.ResponseWriter, *http.Request) {
	h.auseCase.RegisterUser()
}

func GetUserRoutes(r *mux.Router, h *UseCaseHandler) {
	r.HandleFunc("/test", h.Register).Methods("POST")
}
