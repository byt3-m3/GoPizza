package models

import "github.com/google/uuid"

type NewOrderRequest struct {
	CustomerID uuid.UUID
	StoreID    uuid.UUID
}
