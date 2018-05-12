package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/factoryMethod"
)

func TestFactory1(t *testing.T) {

	var Triangle ITriangulo

	f := NewFixture("Test del patrón Factory Method numero 1", t)
	defer f.Run()

	f.Setup(func() {
		Triangle = FabricaTriangulo(4, 2, 3)
	})

	f.Test("Test Triangulo escaleno", func() {
		f.So("...", Triangle.Nombrar(), ShouldEqual, "Soy un triangulo Escaleno")
	})
}


func TestFactory2(t *testing.T) {

	var Triangle ITriangulo

	f := NewFixture("Test del patrón Factory Method numero 2", t)
	defer f.Run()

	f.Setup(func() {
		Triangle = FabricaTriangulo(2, 2, 2)
	})

	f.Test("Test Triangulo equilatero", func() {
		f.So("...", Triangle.Nombrar(), ShouldEqual, "Soy un triangulo Equilatero")
	})
}

func TestFactory3(t *testing.T) {

	var Triangle ITriangulo

	f := NewFixture("Test del patrón Factory Method numero 3", t)
	defer f.Run()

	f.Setup(func() {
		Triangle = FabricaTriangulo(2, 2, 5)
	})

	f.Test("Test Triangulo isoceles", func() {
		f.So("...", Triangle.Nombrar(), ShouldEqual, "Soy un triangulo Isoceles")
	})
}