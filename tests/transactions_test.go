package tests

import (
	"encoding/json"
	"net/http"
	"payment-processor/tests/client"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestCreateTransaction(t *testing.T) {
	accountId, err := setupAccount()
	if err != nil {
		t.Errorf("Error setting up account: %v", err)
		return
	}
	operationTypes := []OperationType{PURCHASE, PURCHASE_WITH_INSTALLMENT_TYPE, WITHDRAW, CREDIT_VOUCHER}
	for _, opType := range operationTypes {
		transaction := Transaction{Account_id: accountId, Operation_type_id: opType, Amount: 10}

		// Marshal the account into JSON format
		payload, err := json.Marshal(transaction)
		if err != nil {
			t.Errorf("Error marshalling account into JSON: %v", err)
			return
		}

		resp, err := client.MakeRequest("POST", "/transactions", &payload)
		if err != nil {
			t.Errorf("Error sending request %v", err)
		}

		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		// Read the response body
		var responseBody Transaction
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
			return
		}

		assert.Equal(t, transaction.Account_id, responseBody.Account_id)
		if opType != CREDIT_VOUCHER {
			assert.Equal(t, transaction.Amount*-1, responseBody.Amount)
		} else {
			assert.Equal(t, transaction.Amount, responseBody.Amount)
		}
		assert.Equal(t, transaction.Operation_type_id, responseBody.Operation_type_id)
	}

}

func TestCreateTransactionWithoutAccountId(t *testing.T) {
	transaction := Transaction{Operation_type_id: 1, Amount: 10}

	// Marshal the account into JSON format
	payload, err := json.Marshal(transaction)
	if err != nil {
		t.Errorf("Error marshalling account into JSON: %v", err)
		return
	}

	resp, err := client.MakeRequest("POST", "/transactions", &payload)
	if err != nil {
		t.Errorf("Error sending request %v", err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}

func TestCreateTransactionWithoutOperationTypeId(t *testing.T) {
	accountId, err := setupAccount()

	if err != nil {
		t.Errorf("Error setting up account: %v", err)
		return
	}

	transaction := Transaction{Account_id: accountId, Amount: 10}

	// Marshal the account into JSON format
	payload, err := json.Marshal(transaction)
	if err != nil {
		t.Errorf("Error marshalling account into JSON: %v", err)
		return
	}

	resp, err := client.MakeRequest("POST", "/transactions", &payload)
	if err != nil {
		t.Errorf("Error sending request %v", err)
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}

func TestCreateTransactionWithoutAmount(t *testing.T) {
	accountId, err := setupAccount()

	if err != nil {
		t.Errorf("Error setting up account: %v", err)
		return
	}

	transaction := Transaction{Account_id: accountId, Operation_type_id: 1}

	// Marshal the account into JSON format
	payload, err := json.Marshal(transaction)
	if err != nil {
		t.Errorf("Error marshalling account into JSON: %v", err)
		return
	}

	resp, err := client.MakeRequest("POST", "/transactions", &payload)
	if err != nil {
		t.Errorf("Error sending request %v", err)
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}
