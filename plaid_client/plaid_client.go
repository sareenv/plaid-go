package plaid_client

import (
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
