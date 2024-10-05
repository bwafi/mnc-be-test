package controllers

import (
	"encoding/json"
	"mnc-golang-test/services"
	"net/http"
)

type CustomerController struct {
	Service *services.CustomerService
}

func (c *CustomerController) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&requestBody)
	customer, err := c.Service.Login(requestBody.Name, requestBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(customer)
}
