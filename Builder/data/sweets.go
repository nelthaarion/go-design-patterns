package sweets

import "fmt"

type ISweetsBuilder interface {
	AddSugar(amount int) ISweetsBuilder
	AddEggs(amount int) ISweetsBuilder
	AddMilk(amount int) ISweetsBuilder
	IsCreamy(isIt bool) ISweetsBuilder
	Bake() Sweets
}

type Sweets interface {
	GetRecipe() string
}

type sweets struct {
	Sugar  int
	Eggs   int
	Milk   int
	Creamy bool
}

func (s *sweets) GetRecipe() string {
	creamy := ""
	if s.Creamy {
		creamy = "creamy"
	} else {
		creamy = "regular"
	}
	return fmt.Sprintf("these delicious sweets made with %d cups of milk, %d eggs and %d spoons of sugar, amazing %s sweets", s.Milk, s.Eggs, s.Sugar, creamy)
}
