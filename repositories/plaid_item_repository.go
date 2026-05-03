package repositories

import (
	"errors"

	"github.com/sareenv/plaid-go/models"
	"gorm.io/gorm"
)

type PlaidItemRepository interface {
	CreatePlaidItem(*models.PlaidItem) error
	FindPlaidItemByItemID(itemID string) (*models.PlaidItem, error)
	FindByUserID(userID uint) ([]models.PlaidItem, error)
}

type plaidItemRepository struct {
	db *gorm.DB
}

func NewPlaidItemRepository(db *gorm.DB) PlaidItemRepository {
	return &plaidItemRepository{db: db}
}

func (pir *plaidItemRepository) CreatePlaidItem(item *models.PlaidItem) error {
	if item == nil {
		return errors.New("item cannot be nil")
	}
	return pir.db.Create(item).Error
}

func (pir *plaidItemRepository) FindPlaidItemByItemID(itemID string) (*models.PlaidItem, error) {
	var item models.PlaidItem
	err := pir.db.Where("item_id = ?", itemID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (pir *plaidItemRepository) FindByUserID(userID uint) ([]models.PlaidItem, error) {
	var items []models.PlaidItem
	err := pir.db.Where("user_id = ?", userID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, err
}
