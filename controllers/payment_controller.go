package controllers

import (
	"encoding/json"
	"mnc-golang-test/services"
	"net/http"
	"strconv"
)

type PaymentController struct {
	Service *services.PaymentService
}

func (c *PaymentController) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		SenderID   string  `json:"sender_id"`
		ReceiverID string  `json:"receiver_id"`
		Amount     float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = c.Service.ProcessPayment(requestBody.SenderID, requestBody.ReceiverID, requestBody.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transfer successful. Amount: " + strconv.FormatFloat(requestBody.Amount, 'f', 2, 64)))
}
