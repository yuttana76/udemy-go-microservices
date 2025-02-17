package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
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
		logs.Error(err.Error())
		return nil, errs.NewInternalServerError()
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
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err.Error())
		return nil, errs.NewInternalServerError()
	}
	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerResponse, nil
}
