package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	cutomerID := mux.Vars(r)["customerID"]

	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("Request body incorredct format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("Request body incorredct format"))
		return
	}

	response, err := h.accSrv.NewAccount(cutomerID, request)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	cutomerID := mux.Vars(r)["customerID"]

	// Get customer from service
	customer, err := h.accSrv.GetAccounts(cutomerID)
	if err != nil {

		handleError(w, err)
		return
	}

	// Return customer as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
