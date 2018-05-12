package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/abstractFactory"
)

func TestAbstract1(t *testing.T) {

	var Factura Factura

	f := NewFixture("Test del patrón Abstract Factory numero 1", t)
	defer f.Run()

	f.Setup(func() {
		Factura = &FacturaCompra{}
	})

	f.Test("Test Factura de compra retorna persona", func() {
		f.So("...", Factura.CrearPersona().Crear(), ShouldEqual, "Soy un Proveedor listo para vender")
	})

	f.Test("Test Factura de compra retorna inventario", func() {
		f.So("...", Factura.CrearInventario().Crear(), ShouldEqual, "Soy un suministro listo para usarse")
	})
}


func TestAbstract2(t *testing.T) {

	var Factura Factura

	f := NewFixture("Test del patrón Abstract Factory numero 2", t)
	defer f.Run()

	f.Setup(func() {
		Factura = &FacturaVenta{}
	})

	f.Test("Test Factura de venta retorna persona", func() {
		f.So("...", Factura.CrearPersona().Crear(), ShouldEqual, "Soy un usuario listo para comprar")
	})

	f.Test("Test Factura de venta retorna inventario", func() {
		f.So("...", Factura.CrearInventario().Crear(), ShouldEqual, "Soy un Producto listo para venderse")
	})
}
