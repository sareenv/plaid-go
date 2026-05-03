package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sareenv/plaid-go/plaidclient"
	"github.com/sareenv/plaid-go/services"
)

type ExchangeHandler struct {
	plaidManager     *plaidclient.PlaidClientManager
	plaidItemService services.PlaidItemService
}

func NewExchangeHandler(plaidManager *plaidclient.PlaidClientManager, pis services.PlaidItemService) *ExchangeHandler {
	return &ExchangeHandler{plaidManager: plaidManager, plaidItemService: pis}
}

type ExchangeTokenRequest struct {
	UserID      string `json:"user_id"`
	PublicToken string `json:"public_token"`
}

type ExchangeTokenResponse struct {
	ItemID  string `json:"item_id"`
	Success bool   `json:"success"`
}

func (eh *ExchangeHandler) ExchangePublicToken(w http.ResponseWriter, r *http.Request) {
	var exchangeTokenRequest ExchangeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&exchangeTokenRequest); err != nil {
		log.Printf("Error decoding request: %v", err)
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "bad_request", Message: "Invalid request body"})
		return
	}

	if exchangeTokenRequest.UserID == "" || exchangeTokenRequest.PublicToken == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "bad_request", Message: "user_id and public_token are required"})
		return
	}

	userID64, err := strconv.ParseUint(exchangeTokenRequest.UserID, 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "invalid_user_id", Message: "Invalid user ID"})
		return
	}

	// Exchange the public token with Plaid for an access token and item ID
	accessToken, itemID, err := eh.plaidManager.ExchangePublicToken(exchangeTokenRequest.PublicToken)
	if err != nil || accessToken == "" || itemID == "" {
		log.Printf("Error exchanging public token: %v", err)
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Code: "exchange_failed", Message: "Could not exchange public token"})
		return
	}

	// Save the user's access token and item ID in the database
	if err := eh.plaidItemService.SavePlaidItem(uint(userID64), accessToken, itemID); err != nil {
		log.Printf("Error saving Plaid item: %v", err)
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Code: "save_plaid_item_failed", Message: "Could not save Plaid item"})
		return
	}

	// Do not return access_token to the client; only confirm success and return item_id
	resp := ExchangeTokenResponse{ItemID: itemID, Success: true}
	writeJSON(w, http.StatusOK, resp)
}
