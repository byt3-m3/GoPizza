package repo

import (
	"GoPizza/models"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

type LocalRepo struct {
	data map[uuid.UUID]*models.Customer
}

func NewLocalCustomerRepo() CustomerRepo {
	data := make(map[uuid.UUID]*models.Customer)
	return LocalRepo{data: data}
}

func (l LocalRepo) Save(customer *models.Customer) error {
	l.data[customer.GetID()] = customer
	return nil

}

func (l LocalRepo) Get(customerId uuid.UUID) (*models.Customer, error) {
	c, ok := l.data[customerId]
	if !ok {
		slog.Warn("customer not found",
			slog.Any("customerID", customerId),
		)
		return nil, fmt.Errorf("customer '%s' not found", customerId.String())
	}

	return c, nil
}
