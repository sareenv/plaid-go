package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/sareenv/plaid-go/plaidclient"
	"github.com/sareenv/plaid-go/services"
)

type LinkHandler struct {
	plaidManager *plaidclient.PlaidClientManager
	userService  services.UserService
}

type LinkTokenRequest struct {
	Email string `json:"email"`
}

type LinkTokenResponse struct {
	LinkToken string `json:"link_token"`
}

func NewLinkHandler(plaidManager *plaidclient.PlaidClientManager, userService services.UserService) *LinkHandler {
	return &LinkHandler{plaidManager: plaidManager, userService: userService}
}

// GetLinkToken handles POST requests to create a Plaid Link token and responds with JSON.
func (lh *LinkHandler) GetLinkToken(w http.ResponseWriter, r *http.Request) {
	var req LinkTokenRequest
	decodingError := json.NewDecoder(r.Body).Decode(&req)
	if decodingError != nil {
		log.Printf("Error decoding request: %v", decodingError)
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Code: "bad_request", Message: "Invalid request body"})
		return
	}
	user, err := lh.userService.GetOrCreateUser(req.Email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Code: "user_creation_failed", Message: "Could not create or fetch user"})
		return
	}
	link, err := lh.plaidManager.CreateLinkToken(strconv.FormatUint(uint64(user.ID), 10))
	if err != nil {
		log.Printf("Error creating link token: %v", err)
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Code: "link_token_creation_failed", Message: "Could not create link token"})
		return
	}
	resp := LinkTokenResponse{LinkToken: link}
	writeJSON(w, http.StatusOK, resp)
}
