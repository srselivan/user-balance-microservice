package handler

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

func (h *Handler) GetBalance() http.Handler {
	request := struct {
		ID int64 `json:"id"`
	}{
		-1,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.Decode(r.Body, &request)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		balance, err := h.service.GetBalance(request.ID)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		bytes, err := json.Encode(balance)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	})
}

func (h *Handler) AppendBalance() http.Handler {
	request := struct {
		ID     int64   `json:"id"`
		Amount float64 `json:"amount"`
	}{
		-1,
		0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.Decode(r.Body, &request)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.AppendBalance(request.ID, request.Amount)
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
