package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/carlos0406/go-gateway/internal/repository"
	"github.com/carlos0406/go-gateway/internal/service"
	"github.com/carlos0406/go-gateway/internal/web/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	if err := godotenv.Load(); err != nil && os.Getenv("USE_ENV_FILE") == "true" {
		log.Fatal("error loading .env file")
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
	)

	log.Println("Connection string:", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database connection: %v", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	invoiceRepository := repository.NewInvoiceRepository(db)
	accountService := service.NewAccountService(accountRepository)
	invoiceService := service.NewInvoiceService(invoiceRepository, *accountService)
	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(accountService, invoiceService, port)
	srv.ConfigureRoutes()
	srv.Start()
}
