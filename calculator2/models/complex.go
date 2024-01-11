package models

import (
	"fmt"
	"strconv"
)

type Complex struct {
	ValidPart     float64
	ImaginaryPart float64
}

func (c *Complex) Add(a Complex) *Complex {
	return &Complex{
		ValidPart:     c.ValidPart + a.ValidPart,
		ImaginaryPart: c.ImaginaryPart + a.ImaginaryPart,
	}
}

func (c *Complex) Sub(a Complex) *Complex {
	return &Complex{
		ValidPart:     c.ValidPart - a.ValidPart,
		ImaginaryPart: c.ImaginaryPart - a.ImaginaryPart,
	}
}

func (c *Complex) Mul(a Complex) *Complex {
	return &Complex{
		ValidPart:     c.ValidPart*a.ValidPart - c.ImaginaryPart*a.ImaginaryPart,
		ImaginaryPart: c.ValidPart*a.ImaginaryPart + c.ImaginaryPart + a.ValidPart,
	}
}

func (c *Complex) Div(a Complex) *Complex {
	return &Complex{
		ValidPart:     (c.ValidPart*a.ValidPart + c.ImaginaryPart*a.ImaginaryPart) / (c.ImaginaryPart*c.ImaginaryPart + a.ImaginaryPart*a.ImaginaryPart),
		ImaginaryPart: (c.ImaginaryPart*a.ValidPart - c.ValidPart*a.ImaginaryPart) / (c.ImaginaryPart*c.ImaginaryPart + a.ImaginaryPart*a.ImaginaryPart),
	}
}

func (c *Complex) ToString() string {
	ValidPart := strconv.FormatFloat(c.ValidPart, 'f', 2, 64)
	ImaginaryPart := strconv.FormatFloat(c.ImaginaryPart, 'f', 2, 64)
	return fmt.Sprintf("%s + %si", ValidPart, ImaginaryPart)
}
