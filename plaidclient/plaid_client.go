package plaidclient

import (
	"context"
	"log"

	"github.com/plaid/plaid-go/v21/plaid"
	"github.com/sareenv/plaid-go/config"
)

type PlaidClientManager struct {
	cfg *config.Config
}

func NewPlaidClient(cfg *config.Config) *PlaidClientManager {
	return &PlaidClientManager{
		cfg: cfg,
	}
}

func (pc *PlaidClientManager) CreatePlaidClient() *plaid.APIClient {
	plaidConfig := plaid.NewConfiguration()
	plaidConfig.AddDefaultHeader("PLAID-CLIENT-ID", pc.cfg.PlaidClientID)
	plaidConfig.AddDefaultHeader("PLAID-SECRET", pc.cfg.PlaidSecret)
	environments := map[string]plaid.Environment{
		"sandbox":     plaid.Sandbox,
		"development": plaid.Development,
		"production":  plaid.Production,
	}
	env, ok := environments[pc.cfg.PlaidEnv]
	if !ok {
		env = plaid.Sandbox
	}
	plaidConfig.UseEnvironment(env)
	client := plaid.NewAPIClient(plaidConfig)
	return client
}

func (pc *PlaidClientManager) ExchangePublicToken(publicToken string) (string, string, error) {
	ctx := context.Background()
	request := plaid.NewItemPublicTokenExchangeRequest(publicToken)
	client := pc.CreatePlaidClient()
	response, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*request).Execute()
	if err != nil {
		return "", "", err
	}
	accessToken := response.GetAccessToken()
	itemId := response.GetItemId()
	return accessToken, itemId, nil
}

func (pc *PlaidClientManager) CreateLinkToken(userID string) (string, error) {
	client := pc.CreatePlaidClient()
	user := plaid.LinkTokenCreateRequestUser{ClientUserId: userID}
	request := plaid.NewLinkTokenCreateRequest(
		"plaid-go", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US, plaid.COUNTRYCODE_CA}, user)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH, plaid.PRODUCTS_TRANSACTIONS})
	resp, _, err := client.PlaidApi.LinkTokenCreate(context.Background()).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		log.Println("Error creating link token: ", err)
		return "", err
	}
	return resp.GetLinkToken(), nil
}
