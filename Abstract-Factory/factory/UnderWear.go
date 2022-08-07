package products

type UnderWear interface {
	SetSize(size int)
	SetColor(color string)
	GetSize() int
	GetColor() string
}

type underWear struct {
	Size  int
	Color string
	For   string
}

func (u *underWear) SetSize(size int) {
	u.Size = size
}
func (u *underWear) SetColor(color string) {
	u.Color = color
}
func (u *underWear) GetSize() int {
	return u.Size
}
func (u *underWear) GetColor() string {
	return u.Color
}

type MenUnderWear struct {
	underWear
}

type WomenUnderWear struct {
	underWear
}
