package service

import (
	"bank/repository"
	"database/sql"
	"errors"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (cs customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := cs.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create new object CustomerResponse
	var customerResponses []CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		})
	}
	return customerResponses, nil
}

func (cs customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := cs.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}
	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerResponse, nil
}
