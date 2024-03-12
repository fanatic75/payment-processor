package transaction

import (
	"encoding/json"
	"net/http"
	"payment-processor/src/core"
	"time"
)

var (
	dbService *core.DbService
)

func InjectDbService(db *core.DbService) {
	dbService = db
}

type OperationType int

const (
	PURCHASE                       OperationType = 1
	PURCHASE_WITH_INSTALLMENT_TYPE OperationType = 2
	WITHDRAW                       OperationType = 3
	CREDIT_VOUCHER                 OperationType = 4
)

type Transaction struct {
	Account_id        int           `json:"account_id" binding:"required"`
	Operation_type_id OperationType `json:"operation_type_id" binding:"required"`
	Amount            float64       `json:"amount" binding:"required"`
	Event_date        time.Time     `json:"event_date"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var transaction Transaction
	decoder.Decode(&transaction)

	if transaction.Account_id == 0 {
		http.Error(w, "account_id is required", http.StatusBadRequest)
		return
	}

	if transaction.Operation_type_id == 0 {
		http.Error(w, "operation_type is required", http.StatusBadRequest)
		return
	}

	if transaction.Amount == 0 {
		http.Error(w, "amount is required", http.StatusBadRequest)
		return
	}

	tx, err := dbService.Client.Begin(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer tx.Rollback(r.Context())

	if transaction.Operation_type_id != CREDIT_VOUCHER {
		transaction.Amount = transaction.Amount * -1
	}

	if err := tx.QueryRow(
		r.Context(),
		"INSERT INTO transactions (account_id, operation_type_id, amount, event_date) VALUES ($1, $2, $3, NOW()) returning account_id, operation_type_id,  amount, event_date", transaction.Account_id, transaction.Operation_type_id, transaction.Amount).Scan(&transaction.Account_id, &transaction.Operation_type_id, &transaction.Amount, &transaction.Event_date); err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	transaction.Event_date = transaction.Event_date.UTC()

	err = tx.Commit(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
