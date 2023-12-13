package models

type Order struct {
}

type Store interface {
	TakeOrder()
}
