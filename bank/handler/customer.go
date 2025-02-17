package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (ch customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	// Get customers from service
	customers, err := ch.custSrv.GetCustomers()
	if err != nil {
		// Handle error
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	// Get customer id from request
	cutomerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// Get customer from service
	customer, err := ch.custSrv.GetCustomer(cutomerID)
	if err != nil {

		handleError(w, err)
		return
	}

	// Return customer as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
