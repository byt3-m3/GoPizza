package models

type DrinkType int

const (
	DrinkTypeCoke   DrinkType = 0
	DrinkTypeSprite DrinkType = 1
)

type Drink interface {
	IsFull() bool
	GetType() DrinkType

	GetPrice() float32
	SetPrice(price float32)
}
