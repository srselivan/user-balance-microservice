package handlers

import (
	"fmt"
	"net/http"
)

func HandleReplenishmentBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler")
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
