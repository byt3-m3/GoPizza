package repo

import (
	"GoPizza/models"
	"github.com/google/uuid"
)

type CustomerRepo interface {
	Save(customer *models.Customer) error
	Get(customerId uuid.UUID) (*models.Customer, error)
}
