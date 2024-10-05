package services

import "errors"

func (s *CustomerService) Logout(customerID string) error {
	customers, err := s.Repo.GetCustomers()
	if err != nil {
		return err
	}

	for i, customer := range customers {
		if customer.ID == customerID && customer.LoggedIn {
			customers[i].LoggedIn = false
			return s.Repo.SaveCustomers(customers)
		}
	}
	return errors.New("customer not found or already logged out")
}
