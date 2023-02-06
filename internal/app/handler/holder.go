package handler

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

func (h *Handler) FreezeAmount() http.Handler {
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
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.FreezeAmount(request.UserID, request.OrderID, request.ServiceID, request.Amount)
		if err != nil {
			logrus.Info(err)
			if err == sql.ErrNoRows {
				NewResponseError(w, err, http.StatusNotFound)
			} else {
				NewResponseError(w, err, http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func (h *Handler) UnFreezeAmount() http.Handler {
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
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.UnFreezeAmount(request.UserID, request.OrderID, request.ServiceID, request.Amount)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
