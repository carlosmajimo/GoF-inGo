package composite

import (
	"fmt"
)

type Combo struct {
	Nombre    string
	Elementos []IProducto
}

func (c *Combo) Vender() (output string){
	output = fmt.Sprintf("El combo %s se esta vendiendo, por eso tambien se venden: ", c.Nombre)
	for _, elemento := range c.Elementos {
		output += elemento.Vender()
	}
	return
}

func (c *Combo) AddProducto(producto IProducto) {
	c.Elementos = append(c.Elementos, producto)
}

func NewCombo(nombre string) *Combo{
	combo := &Combo{
		Nombre: nombre,
	}
	return combo
}