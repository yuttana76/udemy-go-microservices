package repository

import (
	"errors"
	"fmt"
)

type customerRepositoryMock struct {
	customermers []Customer
}

// GetAll implements CustomerRepository.
func (c customerRepositoryMock) GetAll() ([]Customer, error) {
	return c.customermers, nil
}

// GetById implements CustomerRepository.
func (c customerRepositoryMock) GetById(id int) (*Customer, error) {
	idStr := fmt.Sprintf("%03d", id) // Convert id to a zero-padded string
	for _, custolmer := range c.customermers {
		if custolmer.CustomerID == idStr {
			return &custolmer, nil
		}
	}
	return nil, errors.New("customer not found")
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{CustomerID: "001", Name: "Test1", DateOfBirth: "", City: "BKK", ZipCode: "11111", Status: 1},
		{CustomerID: "002", Name: "Test2", DateOfBirth: "", City: "BKK", ZipCode: "22222", Status: 1},
	}

	return customerRepositoryMock{customermers: customers}
}
