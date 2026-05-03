package services

import (
	"github.com/sareenv/plaid-go/models"
	"github.com/sareenv/plaid-go/repositories"
)

type PlaidItemService interface {
	SavePlaidItem(userID uint, accessToken string, itemID string) error
	GetItemsForUser(userID uint) ([]models.PlaidItem, error)
}

type plaidItemService struct {
	repo repositories.PlaidItemRepository
}

func NewPlaidItemService(repo repositories.PlaidItemRepository) PlaidItemService {
	return &plaidItemService{repo: repo}
}

func (pis *plaidItemService) SavePlaidItem(userID uint, accessToken string, itemID string) error {
	item := &models.PlaidItem{
		UserID:      userID,
		ItemID:      itemID,
		AccessToken: accessToken,
	}
	return pis.repo.CreatePlaidItem(item)
}

func (pis *plaidItemService) GetItemsForUser(userID uint) ([]models.PlaidItem, error) {
	return pis.repo.FindByUserID(userID)
}
