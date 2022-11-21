package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/srselivan/user-balance-microservice/internal/database"
)

func readReqBody(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

func HandleAppendBalance() http.Handler {
	reqStruct := struct {
		Id     int64   `json:"id"`
		Amount float64 `json:"amount"`
	}{
		-1,
		0.0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := readReqBody(r, &reqStruct)
		if err != nil {
			logrus.Info(err)
			return
		}

		err = database.AppendBalanceById(reqStruct.Id, reqStruct.Amount)
		if err != nil {
			logrus.Info(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func HandleGetBalance() http.Handler {
	reqStruct := struct {
		Id int64 `json:"id"`
	}{
		-1,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := readReqBody(r, &reqStruct)
		if err != nil {
			logrus.Info(err)
			return
		}

		user, err := database.GetUserById(reqStruct.Id)
		if err != nil {
			logrus.Info(err)
		}

		responseJson, err := json.Marshal(user)
		if err != nil {
			logrus.Info(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	})
}

func HandleTransferBalance() http.Handler {
	reqStruct := struct {
		ReceiveId int64   `json:"receive_id"`
		SendId    int64   `json:"send_id"`
		Amount    float64 `json:"amount"`
	}{
		ReceiveId: -1,
		SendId:    -1,
		Amount:    0.0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := readReqBody(r, &reqStruct)
		if err != nil {
			logrus.Info(err)
			return
		}

		err = database.TransferBalance(reqStruct.SendId, reqStruct.ReceiveId, reqStruct.Amount)
		if err != nil {
			logrus.Info(err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func HandleDebit() http.Handler {
	reqStruct := struct {
		UserId    int64   `json:"user_id"`
		OrderId   int64   `json:"order_id"`
		ServiceId int64   `json:"service_id"`
		Amount    float64 `json:"amount"`
	}{
		UserId:    -1,
		OrderId:   -1,
		ServiceId: -1,
		Amount:    0.0,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := readReqBody(r, &reqStruct)
		if err != nil {
			logrus.Info(err)
			return
		}

		//Connect to db
		//Do work

		w.WriteHeader(http.StatusOK)
	})
}
