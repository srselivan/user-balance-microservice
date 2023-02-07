package handler

import (
	"database/sql"
	"github.com/srselivan/user-balance-microservice/internal/app/model"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/pkg/json"
)

func (h *Handler) FreezeAmount() http.Handler {
	request := model.HolderStruct{
		UserID:    -1,
		OrderID:   -1,
		ServiceID: -1,
		Amount:    0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.Decode(r.Body, &request)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.FreezeAmount(request)
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
	request := model.HolderStruct{
		UserID:    -1,
		OrderID:   -1,
		ServiceID: -1,
		Amount:    0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.Decode(r.Body, &request)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusBadRequest)
			return
		}

		err = h.service.UnFreezeAmount(request)
		if err != nil {
			logrus.Info(err)
			NewResponseError(w, err, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}
