package sweets

type ApplePie struct {
	sugar  int
	eggs   int
	milk   int
	creamy bool
}

func (a *ApplePie) AddSugar(amount int) ISweetsBuilder {
	a.sugar = amount
	return a
}
func (a *ApplePie) AddEggs(amount int) ISweetsBuilder {
	a.eggs = amount
	return a
}
func (a *ApplePie) AddMilk(amount int) ISweetsBuilder {
	a.milk = amount
	return a
}
func (a *ApplePie) IsCreamy(isIt bool) ISweetsBuilder {
	a.creamy = isIt
	return a
}
func (a *ApplePie) Bake() Sweets {
	sweet := &sweets{}
	sweet.Creamy = a.creamy
	sweet.Eggs = a.eggs
	sweet.Milk = a.milk
	sweet.Sugar = a.sugar
	return sweet
}
