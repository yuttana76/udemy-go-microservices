package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {

	query := "INSERT INTO accounts (customer_id, opening_date, account_type,  amount, status ) VALUES(?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, 1)
	if err != nil {
		return nil, err
	}

	// Get account id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	acc.AccountID = int(id)

	return &acc, nil
}

func (r accountRepositoryDB) GetAll(customerId string) ([]Account, error) {

	querys := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE customer_id = ?"
	accounts := []Account{}
	err := r.db.Select(&accounts, querys, customerId)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
