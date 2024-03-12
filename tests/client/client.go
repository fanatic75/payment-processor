package client

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var client *http.Client
var host string

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	host = os.Getenv("TEST_URL") + ":" + os.Getenv("PORT")
	client = &http.Client{}
}

func MakeRequest(method string, url string, payload *[]byte) (*http.Response, error) {
	var data *bytes.Buffer = &bytes.Buffer{}

	if payload != nil {
		data = bytes.NewBuffer(*payload)
	}

	req, err := http.NewRequest(method, host+url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
