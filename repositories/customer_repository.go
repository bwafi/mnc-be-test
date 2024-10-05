package repositories

import (
	"encoding/json"
	"io/ioutil"
	"mnc-golang-test/models"
)

type CustomerRepository struct {
	FilePath string
}

func (r *CustomerRepository) GetCustomers() ([]models.Customer, error) {
	file, err := ioutil.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}
	var customers []models.Customer
	json.Unmarshal(file, &customers)
	return customers, nil
}

func (r *CustomerRepository) SaveCustomers(customers []models.Customer) error {
	data, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.FilePath, data, 0644)
}
