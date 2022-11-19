package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/model"
)

func HandleAppendBalance() http.Handler {
	reqStruct := struct {
		Id     int64   `json:"id"`
		Amount float64 `json:"amount"`
	}{
		0,
		0.0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Info(err)
		}

		err = json.Unmarshal(body, &reqStruct)
		if err != nil {
			logrus.Info(err)
		}

		//ADD AMOUNT TO USER DATA HERE
		w.WriteHeader(http.StatusOK)

	})
}

func HandleGetBalance() http.Handler {
	user := model.NewUser()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logrus.Info(err)
		}

		err = json.Unmarshal(body, &user)
		if err != nil {
			logrus.Info(err)
		}

		//get balance

		user.SetBalance(123) //by db
		responseJson, err := json.Marshal(user)
		if err != nil {
			logrus.Info(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	})
}

func HandleTransferBalance(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}

func HandleDebit(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}
