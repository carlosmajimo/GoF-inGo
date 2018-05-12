package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/adapter"
)

func TestAdapter1(t *testing.T) {

	var Guitarra IGuitar

	f := NewFixture("Test del patrón Adapter numero 1", t)
	defer f.Run()

	f.Setup(func() {
		Guitarra = &ElectricGuitar{}
	})

	f.Test("Test Guitarra electrica encender", func() {
		f.So("...", Guitarra.OnGuitar(), ShouldEqual, "Voy a tocar la guitarra electrica")
	})

	f.Test("Test Guitarra electrica apagar", func() {
		f.So("...", Guitarra.OffGuitar(), ShouldEqual, "Estoy cansado de tocar la guitarra electrica")
	})
}

func TestAdapter2(t *testing.T) {

	var Guitarra IGuitar

	f := NewFixture("Test del patrón Adapter numero 2", t)
	defer f.Run()

	f.Setup(func() {
		Guitarra = &AcousticGuitarAdapter{}
	})

	f.Test("Test Guitarra acustica tocar", func() {
		f.So("...", Guitarra.OnGuitar(), ShouldEqual, "Voy a tocar la guitarra acustica")
	})

	f.Test("Test Guitarra acustica dejar de tocar", func() {
		f.So("...", Guitarra.OffGuitar(), ShouldEqual, "Estoy cansado de tocar la guitarra acustica")
	})
}
