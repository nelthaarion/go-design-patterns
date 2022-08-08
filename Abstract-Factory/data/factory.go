package products

import (
	"errors"
	"strings"
)

type GenderFactory interface {
	MakeJeans() Jeans
	MakeTShirt() TShirt
	MakeUnderWear() UnderWear
}

type Man struct{}

func (m *Man) MakeJeans() Jeans {
	return &MenJeans{
		jeans{
			Size:  34,
			Color: "blue",
			For:   "MEN",
		},
	}
}
func (m *Man) MakeTShirt() TShirt {
	return &MenTShirt{
		Tshirt{
			Size:  30,
			Color: "blac",
			For:   "MEN",
		},
	}
}
func (m *Man) MakeUnderWear() UnderWear {
	return &MenUnderWear{
		underWear{
			Color: "gray",
			Size:  34,
			For:   "men",
		},
	}
}

type Woman struct{}

func (w *Woman) MakeJeans() Jeans {
	return &WomenJeans{
		jeans{
			Size:  24,
			Color: "blue",
			For:   "WOMEN",
		},
	}
}
func (w *Woman) MakeTShirt() TShirt {
	return &WomenTShirt{
		Tshirt{
			Size:  20,
			Color: "blac",
			For:   "WOMEN",
		},
	}
}
func (w *Woman) MakeUnderWear() UnderWear {
	return &WomenUnderWear{
		underWear{
			Color: "gray",
			Size:  24,
			For:   "WOMEN",
		},
	}
}

func ForGender(gender string) (GenderFactory, error) {
	if strings.ToLower(gender) == "man" {
		return &Man{}, nil
	} else if strings.ToLower(gender) == "woman" {
		return &Woman{}, nil
	}
	return nil, errors.New("bad input")
}
