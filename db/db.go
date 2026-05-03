package db

import (
	"github.com/sareenv/plaid-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	config *config.Config
}

func NewDBManager(cfg *config.Config) *DBManager {
	return &DBManager{
		config: cfg,
	}
}

func (dbm *DBManager) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dbm.config.DatabaseURL), &gorm.Config{})
}
