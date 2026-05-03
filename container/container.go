package container

import (
	"errors"
	"log"

	"github.com/sareenv/plaid-go/config"
	"github.com/sareenv/plaid-go/db"
	"github.com/sareenv/plaid-go/plaidclient"
	"github.com/sareenv/plaid-go/repositories"
	"github.com/sareenv/plaid-go/services"
	"gorm.io/gorm"
)

// Container owns application-wide dependencies and wires them together.
type Container struct {
	cfg *config.Config
	db  *gorm.DB

	// singletons (lazy init)
	plaidManager     *plaidclient.PlaidClientManager
	userRepo         repositories.UserRepository
	plaidItemRepo    repositories.PlaidItemRepository
	userService      services.UserService
	plaidItemService services.PlaidItemService
}

// New creates a new Container by loading config and establishing a DB connection.
func New() (*Container, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dbManager := db.NewDBManager(cfg)
	if dbManager == nil {
		return nil, errors.New("failed to initialize DBManager")
	}
	conn, err := dbManager.Connect()
	if err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, errors.New("failed to get DB connection")
	}
	log.Println("Connected to database")

	return &Container{cfg: cfg, db: conn}, nil
}

func (c *Container) Config() *config.Config { return c.cfg }
func (c *Container) DB() *gorm.DB           { return c.db }

// PlaidManager returns a singleton PlaidClientManager.
func (c *Container) PlaidManager() *plaidclient.PlaidClientManager {
	if c.plaidManager == nil {
		c.plaidManager = plaidclient.NewPlaidClient(c.cfg)
	}
	return c.plaidManager
}

// UserRepository returns a singleton UserRepository.
func (c *Container) UserRepository() repositories.UserRepository {
	if c.userRepo == nil {
		c.userRepo = repositories.NewUserRepository(c.db)
	}
	return c.userRepo
}

// PlaidItemRepository returns a singleton PlaidItemRepository.
func (c *Container) PlaidItemRepository() repositories.PlaidItemRepository {
	if c.plaidItemRepo == nil {
		c.plaidItemRepo = repositories.NewPlaidItemRepository(c.db)
	}
	return c.plaidItemRepo
}

// UserService returns a singleton UserService.
func (c *Container) UserService() services.UserService {
	if c.userService == nil {
		c.userService = services.NewUserService(c.UserRepository())
	}
	return c.userService
}

// PlaidItemService returns a singleton PlaidItemService.
func (c *Container) PlaidItemService() services.PlaidItemService {
	if c.plaidItemService == nil {
		c.plaidItemService = services.NewPlaidItemService(c.PlaidItemRepository())
	}
	return c.plaidItemService
}
