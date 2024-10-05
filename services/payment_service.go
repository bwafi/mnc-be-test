package services

import (
	"errors"
	"mnc-golang-test/models"
	"mnc-golang-test/repositories"
	"time"
)

type PaymentService struct {
	CustomerRepo *repositories.CustomerRepository
	HistoryRepo  *repositories.HistoryRepository
}

func (s *PaymentService) ProcessPayment(senderID, receiverID string, amount float64) error {
	customers, err := s.CustomerRepo.GetCustomers()
	if err != nil {
		return err
	}

	var sender, receiver *models.Customer

	for i, customer := range customers {
		if customer.ID == senderID {
			sender = &customers[i]
		}
		if customer.ID == receiverID {
			receiver = &customers[i]
		}
	}

	if sender == nil || receiver == nil {
		return errors.New("sender or receiver not found")
	}

	if !sender.LoggedIn {
		return errors.New("sender is not logged in")
	}

	if sender.Balance < amount {
		return errors.New("insufficient balance")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	err = s.CustomerRepo.SaveCustomers(customers)
	if err != nil {
		return err
	}

	history := models.History{
		CustomerID: senderID,
		Action:     "Transfer",
		Amount:     amount,
		Timestamp:  time.Now(),
		Details:    "Transfer to customer ID " + receiverID,
	}
	return s.HistoryRepo.AddHistory(history)
}
