package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/composite"
)

func TestComposite(t *testing.T) {

	var medicina IProducto
	var comida IProducto
	var combo *Combo

	f := NewFixture("Test del patrón Composite", t)
	defer f.Run()

	f.Setup(func() {
		medicina = NewMedicamento("Acetaminofen", 4)
		comida = NewComida("pan", "perecedero")
		combo = NewCombo("PanDolor")
		combo.AddProducto(comida)
		combo.AddProducto(comida)
	})

	f.Test("Test vender medicamento", func() {
		f.So("...", medicina.Vender(), ShouldEqual, " Se vendió el medicamento Acetaminofen en la dosis 4.")
	})

	f.Test("Test vender comida", func() {
		f.So("...", comida.Vender(), ShouldEqual, " Se vendió la comida pan y es de tipo perecedero.")
	})

	f.Test("Test vender combo", func() {
		expected := "El combo PanDolor se esta vendiendo, por eso tambien se venden:  Se vendió la comida pan y es de tipo perecedero. Se vendió la comida pan y es de tipo perecedero."
		f.So("...", combo.Vender(), ShouldEqual, expected)
	})

}