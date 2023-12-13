package models

import "GoPizza/models"

type CreateCustomerRequest struct {
	Name            string
	StartingBalance float32
}

type CustomerResponse struct {
	*models.Customer
	Operations []Operation
}
