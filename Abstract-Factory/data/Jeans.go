package products

type Jeans interface {
	SetSize(size int)
	SetColor(color string)
	GetSize() int
	GetColor() string
}

type jeans struct {
	Size  int
	Color string
	For   string
}

func (j *jeans) SetSize(size int) {
	j.Size = size
}
func (j *jeans) SetColor(color string) {
	j.Color = color
}

func (j *jeans) GetSize() int {
	return j.Size
}
func (j *jeans) GetColor() string {
	return j.Color
}

type MenJeans struct {
	jeans
}

type WomenJeans struct {
	jeans
}
