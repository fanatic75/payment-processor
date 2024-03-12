package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment-processor/tests/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Account struct {
	Id             int `json:"account_id"`
	DocumentNumber int `json:"document_number"`
}

func TestCreateAccount(t *testing.T) {

	// Create an instance of the Account struct
	account := Account{DocumentNumber: 123456789}

	// Marshal the account into JSON format
	payload, err := json.Marshal(account)
	if err != nil {
		t.Errorf("Error marshalling account into JSON: %v", err)
		return
	}

	resp, err := client.MakeRequest("POST", "/accounts", &payload)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Read the response body
	var responseBody Account
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
		return
	}

	assert.Equal(t, account.DocumentNumber, responseBody.DocumentNumber)
}

func TestCreateAccountWithoutDocumentNumber(t *testing.T) {

	// Create an instance of the Account struct
	account := Account{DocumentNumber: 0}

	// Marshal the account into JSON format
	payload, err := json.Marshal(account)
	if err != nil {
		t.Errorf("Error marshalling account into JSON: %v", err)
		return
	}

	resp, err := client.MakeRequest("POST", "/accounts", &payload)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetAccount(t *testing.T) {
	accountId, err := setupAccount()
	if err != nil {
		t.Errorf("Error setting up account: %v", err)
		return
	}

	resp, err := client.MakeRequest("GET", "/accounts/"+fmt.Sprint(accountId), nil)

	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	var responseBody Account
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
		return
	}

	assert.Equal(t, accountId, responseBody.Id)
}

func TestGetAccountNotFound(t *testing.T) {
	// Create a new request
	resp, err := client.MakeRequest("GET", "/accounts/sdfdf", nil)
	if err != nil {
		t.Errorf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func setupAccount() (int, error) {
	account := Account{DocumentNumber: 123456789}
	payload, err := json.Marshal(account)
	if err != nil {
		return 0, err
	}
	resp, err := client.MakeRequest("POST", "/accounts", &payload)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	var responseBody Account
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return 0, err
	}

	return responseBody.Id, nil

}
