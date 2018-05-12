package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "Gof-inGo/project/singleton"
)

func TestSingleton1(t *testing.T) {

	singleton1 := Person()
	singleton2 := Person()

	f := NewFixture("Test del patrón Singleton numero 1", t)
	defer f.Run()

	f.Test("Test variables diferentes misma instancia", func() {
		singleton1.SetNombre("Carlos")
		singleton2.SetNombre("Mario")
		f.So("...", singleton1.GetNombre(), ShouldEqual, singleton2.GetNombre())
	})
}


func TestSingleton2(t *testing.T) {

	singleton1 := Person()
	singleton2 := Person()

	f := NewFixture("Test del patrón Singleton numero 2", t)
	defer f.Run()

	f.Test("Test variable variable 1 con nombre diferente despues de la segunda instancia", func() {
		singleton1.SetNombre("Carlos")
		singleton2.SetNombre("Mario")
		f.So("...", singleton1.GetNombre(), ShouldNotEqual, "Carlos")
	})
}
