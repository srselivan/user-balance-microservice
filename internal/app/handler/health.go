package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

func (h *Handler) Health() http.Handler {
	healthAnswer := struct {
		Message string `json:"message"`
		Status  string `json:"status"`
		Code    int    `json:"code"`
	}{
		fmt.Sprintf("Responce from health check at %v", time.Now()),
		http.StatusText(http.StatusOK),
		http.StatusOK,
	}

	responseJson, _ := json.Encode(healthAnswer)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	})
}
