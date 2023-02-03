package handlers

import (
	"net/http"

	"github.com/srselivan/user-balance-microservice/internal/app/store"
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

type Handler struct {
	store *store.Store
}

func New(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) createResponseError(w http.ResponseWriter, error error, status int) {
	w.WriteHeader(status)
	json, _ := json.Encode(error.Error())
	w.Write(json)
}
