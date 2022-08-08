package sweets

type Pancake struct {
	sugar  int
	eggs   int
	milk   int
	creamy bool
}

func (p *Pancake) AddSugar(amount int) ISweetsBuilder {
	p.sugar = amount
	return p
}
func (p *Pancake) AddEggs(amount int) ISweetsBuilder {
	p.eggs = amount
	return p
}
func (p *Pancake) AddMilk(amount int) ISweetsBuilder {
	p.milk = amount
	return p
}
func (p *Pancake) IsCreamy(isIt bool) ISweetsBuilder {
	p.creamy = isIt
	return p
}
func (p *Pancake) Bake() Sweets {
	sweet := &sweets{}
	sweet.Creamy = p.creamy
	sweet.Eggs = p.eggs
	sweet.Milk = p.milk
	return sweet
}
