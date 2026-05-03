package container

import (
    "github.com/sareenv/plaid-go/handlers"
)

// NewExchangeHandler constructs an ExchangeHandler using container-managed dependencies.
func (c *Container) NewExchangeHandler() *handlers.ExchangeHandler {
    return handlers.NewExchangeHandler(c.PlaidManager(), c.PlaidItemService())
}
