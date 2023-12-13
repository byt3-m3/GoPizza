package models

type Consumable interface {
	Consume() error
}
