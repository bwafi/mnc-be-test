package main

import (
	"mnc-golang-test/controllers"
	"mnc-golang-test/repositories"
	"mnc-golang-test/services"
	"net/http"
)

func main() {
	customerRepo := &repositories.CustomerRepository{FilePath: "data/customers.json"}
	historyRepo := &repositories.HistoryRepository{FilePath: "data/history.json"}

	customerService := &services.CustomerService{Repo: customerRepo}
	paymentService := &services.PaymentService{CustomerRepo: customerRepo, HistoryRepo: historyRepo}

	customerController := &controllers.CustomerController{Service: customerService}
	logoutController := &controllers.LogoutController{Service: customerService}
	paymentController := &controllers.PaymentController{Service: paymentService}

	http.HandleFunc("/login", customerController.Login)
	http.HandleFunc("/logout", logoutController.Logout)
	http.HandleFunc("/payment", paymentController.ProcessPayment)

	http.ListenAndServe(":8080", nil)
}
