package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"payment-processor/src/accounts"
	"payment-processor/src/core"
	"payment-processor/src/transaction"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
)

var (
	dbService *core.DbService
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// List of required environment variables
	requiredEnvVars := []string{"PORT", "DATABASE_URI"}

	// Check if all required environment variables are present
	for _, envVar := range requiredEnvVars {
		value, exists := os.LookupEnv(envVar)
		if !exists || value == "" {
			log.Fatalf("Error: Required environment variable %s is missing or empty", envVar)
		}
	}

	// Connect to the database
	dbService = &core.DbService{}
	err := dbService.InitDbClient(os.Getenv("DATABASE_URI"), context.Background())
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func main() {
	// Inject the database service into the accounts and transaction packages
	accounts.InjectDbService(dbService)
	transaction.InjectDbService(dbService)

	r := chi.NewRouter()
	r.Use()

	r.Mount("/accounts", accounts.SetupRoutes())
	r.Mount("/transactions", transaction.SetupRoutes())

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
