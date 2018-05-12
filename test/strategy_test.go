package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/strategy"
)

func TestStrategy1(t *testing.T) {

	var vehiculo *Vehiculo

	f := NewFixture("Test del patrón Strategy numero 1", t)
	defer f.Run()

	f.Setup(func() {
		vehiculo = &Vehiculo{nil}
		vehiculo.ConduccionNormal()
	})

	f.Test("Test conduccion normal", func() {
		f.So("...", vehiculo.Acelerar(10), ShouldEqual, "Tipo de conducción: Conduccion Normal, con la potencia de 13 se aumento la velocidad en 7")
	})
}

func TestStrategy2(t *testing.T) {

	var vehiculo *Vehiculo

	f := NewFixture("Test del patrón Sstrategy numero 2", t)
	defer f.Run()

	f.Setup(func() {
		vehiculo = &Vehiculo{nil}
		vehiculo.ConduccionNormal()
	})

	f.Test("Test conduccion deportiva", func() {
		f.So("...", vehiculo.Acelerar(10), ShouldEqual, "Tipo de conducción: Conduccion Normal, con la potencia de 13 se aumento la velocidad en 7")
	})
}