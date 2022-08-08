package products

type TShirt interface {
	SetSize(size int)
	SetColor(color string)
	GetSize() int
	GetColor() string
}

type Tshirt struct {
	Size  int
	Color string
	For   string
}

func (t *Tshirt) SetSize(size int) {
	t.Size = size
}
func (t *Tshirt) SetColor(color string) {
	t.Color = color
}

func (t *Tshirt) GetSize() int {
	return t.Size
}
func (t *Tshirt) GetColor() string {
	return t.Color
}

type MenTShirt struct {
	Tshirt
}

type WomenTShirt struct {
	Tshirt
}
