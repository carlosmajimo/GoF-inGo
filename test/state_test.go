package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/state"
)

func TestState1(t *testing.T) {

	var tiempo *Tiempo

	f := NewFixture("Test del patrón State numero 1", t)
	defer f.Run()

	f.Setup(func() {
		tiempo = &Tiempo{}
		tiempo.SetContexto(&Lluvioso{})
	})

	f.Test("Test tiempo lluvioso", func() {
		f.So("...", tiempo.Request(), ShouldEqual, "Está lloviendo")
	})
}

func TestState2(t *testing.T) {

	var tiempo *Tiempo

	f := NewFixture("Test del patrón State numero 2", t)
	defer f.Run()

	f.Setup(func() {
		tiempo = &Tiempo{}
		tiempo.SetContexto(&Nublado{})
	})

	f.Test("Test tiempo nublado", func() {
		f.So("...", tiempo.Request(), ShouldEqual, "Es un dia nublado")
	})
}

func TestState3(t *testing.T) {

	var tiempo *Tiempo

	f := NewFixture("Test del patrón State numero 3", t)
	defer f.Run()

	f.Setup(func() {
		tiempo = &Tiempo{}
		tiempo.SetContexto(&Nevado{})
	})

	f.Test("Test tiempo neblado", func() {
		f.So("...", tiempo.Request(), ShouldEqual, "Está Nevando")
	})
}

func TestState4(t *testing.T) {

	var tiempo *Tiempo

	f := NewFixture("Test del patrón State numero 4", t)
	defer f.Run()

	f.Setup(func() {
		tiempo = &Tiempo{}
		tiempo.SetContexto(&Soleado{})
	})

	f.Test("Test tiempo soleado", func() {
		f.So("...", tiempo.Request(), ShouldEqual, "Es un dia soleado")
	})
}