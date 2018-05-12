package composite

import (
	"fmt"
)

type Medicamento struct {
	Nombre string
	Dosis  int
}

func (c *Medicamento) Vender() (output string){
	output = fmt.Sprintf(" Se vendió el medicamento %s en la dosis %d.", c.Nombre, c.Dosis)
	return
}

func NewMedicamento(nombre string, dosis int) *Medicamento{
	combo := &Medicamento{
		Nombre: nombre,
		Dosis: dosis,
	}
	return combo
}