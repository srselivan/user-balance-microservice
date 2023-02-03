package handlers

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

func (h *Handler) ReserveMoney() http.Handler {
	request := struct {
		UserID    int64   `json:"user_id"`
		OrderID   int64   `json:"order_id"`
		ServiceID int64   `json:"service_id"`
		Amount    float64 `json:"amount"`
	}{
		-1,
		-1,
		-1,
		-1,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.Decode(r.Body, &request)
		if err != nil {
			logrus.Info(err)
			h.createResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.store.ReserveMoney(request.UserID, request.OrderID, request.ServiceID, request.Amount)
		if err != nil {
			logrus.Info(err)
			if err == sql.ErrNoRows {
				h.createResponseError(w, err, http.StatusNotFound)
			} else {
				h.createResponseError(w, err, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
