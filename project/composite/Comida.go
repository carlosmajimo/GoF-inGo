package composite

import (
	"fmt"
)

type Comida struct {
	Nombre string
	Tipo   string
}

func (c *Comida) Vender() (output string){
	output = fmt.Sprintf(" Se vendió la comida %s y es de tipo %s.", c.Nombre, c.Tipo)
	return
}

func NewComida(nombre string, tipo string) *Comida{
	combo := &Comida{
		Nombre: nombre,
		Tipo: tipo,
	}
	return combo
}
