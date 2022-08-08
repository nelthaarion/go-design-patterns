package sweets

import (
	"errors"
	"strings"
)

type IBaker interface {
	GetToBake() Sweets
}

type Baker struct {
}

func (b *Baker) GetToBake(sweets string) (ISweetsBuilder, error) {
	if strings.ToLower(sweets) == "pancake" {
		return &Pancake{}, nil
	} else if strings.ToLower(sweets) == "applepie" {
		return &ApplePie{}, nil
	} else {
		return nil, errors.New("bad input")
	}
}
