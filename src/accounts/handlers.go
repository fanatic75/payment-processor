package accounts

import (
	"encoding/json"
	"net/http"
	"payment-processor/src/core"

	"github.com/go-chi/chi/v5"
)

var (
	dbService *core.DbService
)

func InjectDbService(db *core.DbService) {
	dbService = db
}

type Account struct {
	Id              int `json:"account_id"`
	Document_number int `json:"document_number" binding:"required"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	decorder := json.NewDecoder(r.Body)
	var account Account

	if err := decorder.Decode(&account); err != nil {
		http.Error(w, "Bad payload", http.StatusBadRequest)
		return
	}

	if account.Document_number == 0 {
		http.Error(w, "document_number is required", http.StatusBadRequest)
		return
	}

	if err := dbService.Client.QueryRow(r.Context(), "INSERT INTO accounts (document_number) VALUES ($1) RETURNING id, document_number", account.Document_number).Scan(&account.Id, &account.Document_number); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var account Account

	if err := dbService.Client.QueryRow(r.Context(), "SELECT id, document_number FROM accounts WHERE id=$1", id).Scan(&account.Id, &account.Document_number); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
