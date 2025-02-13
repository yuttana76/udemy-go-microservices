package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {

	// Create db instance
	db, err := sqlx.Open("mysql", "root:password102@tcp(localhost:3306)/my-database")
	if err != nil {
		panic(err)
	}

	// Create customer instance
	customerRepository := repository.NewCustomerRepositoryDB(db)

	// Test repository
	// customer, err := customerRepository.GetById(001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)

	// customers, err := customerRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	// Test service
	customerService := service.NewCustomerService(customerRepository)
	// customers, err := customerService.GetCustomer(002)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customers)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)

}
