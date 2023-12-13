package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPizzaFactory(t *testing.T) {
	type args struct {
		pizzaType  PizzaType
		pizzaPrice float32
	}

	type testCase struct {
		name          string
		testArgs      args
		expectedPizza Pizza
	}

	testCases := []testCase{
		{
			testArgs: args{
				pizzaType:  PizzaTypeCheese,
				pizzaPrice: 1.50,
			},
			expectedPizza: &pizza{
				price:     1.50,
				pizzaType: PizzaTypeCheese,
			},
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			p := NewPizza(NewPizzaInput{
				PizzaType: tc.testArgs.pizzaType,
				Price:     tc.testArgs.pizzaPrice,
			})

			assert.Equal(t, p.GetPrice(), tc.expectedPizza.GetPrice())
			assert.Equal(t, p.GetType(), tc.expectedPizza.GetType())
		})
	}

}
