package handlers

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
			h.createResponseError(w, err, http.StatusBadRequest)
			return
		}

		user, err := h.store.GetUserById(request.ID)
		if err != nil {
			logrus.Info(err)
			if err == sql.ErrNoRows {
				h.createResponseError(w, err, http.StatusNotFound)
			} else {
				h.createResponseError(w, err, http.StatusInternalServerError)
			}
			return
		}

		json, err := json.Encode(user.Balance)
		if err != nil {
			logrus.Info(err)
			h.createResponseError(w, err, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
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
			h.createResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.store.AppendBalanceByUserId(request.ID, request.Amount)
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
