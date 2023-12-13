package models

import "fmt"

type PizzaType int

const (
	PizzaTypeCheese    PizzaType = 0
	PizzaTypePepperoni PizzaType = 1
)

type Pizza interface {
	HasSauce() bool
	HasCheese() bool
	GetType() PizzaType

	GetPrice() float32
	SetPrice(price float32)
}

type NewPizzaInput struct {
	PizzaType PizzaType
	Price     float32
}

func NewPizza(input NewPizzaInput) Pizza {

	return &pizza{
		price:     input.Price,
		pizzaType: input.PizzaType,
		lifeBar:   100,
	}

}

type pizza struct {
	price     float32
	pizzaType PizzaType
	lifeBar   int
}

func (p *pizza) Consume() error {
	if p.lifeBar > 0 {
		p.lifeBar = -25
	}

	if p.lifeBar < 0 {
		return fmt.Errorf("pizza")
	}

	return nil

}

func (p *pizza) HasSauce() bool {
	return true
}

func (p *pizza) HasCheese() bool {
	return true
}

func (p *pizza) GetType() PizzaType {
	return p.pizzaType
}

func (p *pizza) GetPrice() float32 {
	return p.price
}

func (p *pizza) SetPrice(price float32) {
	p.price = price

}
