package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/sareenv/plaid-go/config"
	"github.com/sareenv/plaid-go/db"
	"gorm.io/gorm"
)

type Response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encodeErr := json.NewEncoder(w).Encode(Response{
			Status: "success",
			Body:   "database connected",
		})
		if encodeErr != nil {
			log.Printf("failed to encode success response: %v", encodeErr)
		}
	}
}

func setupServer(port string) {
	fmt.Printf("Starting server on port %s\n", port)
	http.HandleFunc("/", rootHandler())
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("Failed to start server: %v", err)
		return
	}
}

func connectDB(cfg *config.Config) (*gorm.DB, error) {
	dbManager := db.NewDBManager(cfg)
	if dbManager == nil {
		log.Printf("Failed to initialize DBManager")
		return nil, errors.New("failed to initialize DBManager")
	}
	connection, connError := dbManager.Connect()
	if connError != nil {
		log.Printf("Failed to connect to database:%v", connError)
		return nil, connError
	} else if connection == nil {
		log.Println("Failed to get connection")
		return nil, errors.New("failed to get connection")
	}
	return connection, nil
}

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatalf("Error loading config: %v", cfgErr)
	}
	port := fmt.Sprintf(":%s", cfg.Port)
	_, err := connectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	setupServer(port)
}
