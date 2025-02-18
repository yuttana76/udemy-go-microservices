package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerId string, request NewAccountRequest) (*AccountResponse, error) {
	// Validate
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("Amount must be at least 5000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("Account type must be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}

	response := AccountResponse{AccountID: newAcc.AccountID, CustomerID: newAcc.CustomerID, OpeningDate: newAcc.OpeningDate, AccountType: newAcc.AccountType, Amount: newAcc.Amount, Status: newAcc.Status}

	return &response, err

}

func (s accountService) GetAccounts(customerId string) ([]AccountResponse, error) {

	accounts, err := s.accRepo.GetAll(customerId)
	if err != nil {
		logs.Error((err))
		return nil, errs.NewInternalServerError()
	}

	response := []AccountResponse{}
	for _, a := range accounts {
		response = append(response, AccountResponse{AccountID: a.AccountID, CustomerID: a.CustomerID, OpeningDate: a.OpeningDate, AccountType: a.AccountType, Amount: a.Amount, Status: a.Status})
	}
	return response, nil
}
