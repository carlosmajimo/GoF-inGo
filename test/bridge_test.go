package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/bridge"
)

func TestBridge1(t *testing.T) {

	var circulo *Circulo

	f := NewFixture("Test del patrón Bridge numero 1", t)
	defer f.Run()

	f.Setup(func() {
		circulo = &Circulo{
			DrawingContext: &OpenGL{},
			Center: Point{X: 100, Y: 100},
			Radius: 50,
		}
	})

	f.Test("Test Dibujar circulo con OpenGL", func() {
		f.So("...", circulo.Draw(), ShouldEqual, "OpenGL está dibujando el elipse en el rectangulo {{50 50} {100 100}}")
	})
}

func TestBridge2(t *testing.T) {

	var circulo *Circulo

	f := NewFixture("Test del patrón Bridge numero 2", t)
	defer f.Run()

	f.Setup(func() {
		circulo = &Circulo{
			DrawingContext: &Direct2D{},
			Center: Point{X: 100, Y: 100},
			Radius: 50,
		}
	})

	f.Test("Test Dibujar circulo con Direct2D", func() {
		f.So("...", circulo.Draw(), ShouldEqual, "Direct2D está dibujando el elipse en el rectangulo {{50 50} {100 100}}")
	})
}
