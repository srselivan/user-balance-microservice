package handlers

import (
	"fmt"
	"net/http"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func HandleReplenishmentBalance(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}

func HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}

func HandleTransferBalance(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}

func HandleDebit(w http.ResponseWriter, r *http.Request) {
	//Replenishment of the balance
}
