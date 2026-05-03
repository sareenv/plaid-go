package container

import (
    "github.com/sareenv/plaid-go/handlers"
)

// NewLinkHandler constructs a LinkHandler using container-managed dependencies.
func (c *Container) NewLinkHandler() *handlers.LinkHandler {
    return handlers.NewLinkHandler(c.PlaidManager(), c.UserService())
}
