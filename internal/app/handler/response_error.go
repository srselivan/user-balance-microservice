package handler

import (
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
	"net/http"
)

func NewResponseError(w http.ResponseWriter, error error, statusCode int) {
	w.WriteHeader(statusCode)
	bytes, _ := json.Encode(error.Error())
	w.Write(bytes)
}
