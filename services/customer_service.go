package services

import (
	"errors"
	"mnc-golang-test/models"
	"mnc-golang-test/repositories"
)

type CustomerService struct {
	Repo *repositories.CustomerRepository
}

func (s *CustomerService) Login(name, password string) (*models.Customer, error) {
	customers, _ := s.Repo.GetCustomers()

	for i, customer := range customers {
		if customer.Name == name && customer.Password == password {
			customers[i].LoggedIn = true
			s.Repo.SaveCustomers(customers)
			return &customers[i], nil
		}
	}
	return nil, errors.New("invalid credentials")
}
