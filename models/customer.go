package models

import "github.com/google/uuid"

//type Customer interface {
//	GetID() uuid.UUID
//	VerifyFunds(amount float32) bool
//}

type Customer struct {
	ID             uuid.UUID
	Name           string
	AccountBalance float32
}

type NewCustomerInput struct {
	ID               uuid.UUID
	Name             string
	BeginningBalance float32
}

func NewCustomer(input NewCustomerInput) *Customer {

	return &Customer{
		ID:             input.ID,
		Name:           input.Name,
		AccountBalance: input.BeginningBalance,
	}
}

func (c Customer) GetID() uuid.UUID {
	return c.ID
}

func (c Customer) VerifyFunds(amount float32) bool {
	if c.AccountBalance > amount {
		return true
	}

	return false
}
