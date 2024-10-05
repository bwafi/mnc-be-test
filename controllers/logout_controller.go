package controllers

import (
	"encoding/json"
	"mnc-golang-test/services"
	"net/http"
)

type LogoutController struct {
	Service *services.CustomerService
}

func (c *LogoutController) Logout(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		CustomerID string `json:"customer_id"`
	}
	json.NewDecoder(r.Body).Decode(&requestBody)

	err := c.Service.Logout(requestBody.CustomerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
