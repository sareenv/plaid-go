package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sareenv/plaid-go/container"
)

func setupServer(port string) {
	fmt.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Printf("Failed to start server: %v", err)
		return
	}
}

func main() {
	c, err := container.New()
	if err != nil {
		log.Fatalf("failed to initialize container: %v", err)
	}

	port := fmt.Sprintf(":%s", c.Config().Port)

	// Build handlers via factories
	linkHandler := c.NewLinkHandler()
	exchangeHandler := c.NewExchangeHandler()

	// Routes
	http.HandleFunc("/api/link/token", linkHandler.GetLinkToken)
	http.HandleFunc("/api/link/exchange", exchangeHandler.ExchangePublicToken)

	setupServer(port)
}
