package abstractFactory

type Factura interface {
	CrearInventario() IInventario
	CrearPersona() IPersona
}

type FacturaCompra struct{}

func (f *FacturaCompra) CrearInventario() IInventario {
	return &Suministro{}
}

func (f *FacturaCompra) CrearPersona() IPersona {
	return &Proveedor{}
}

type FacturaVenta struct{}

func (f *FacturaVenta) CrearInventario() IInventario {
	return &Producto{}
}

func (f *FacturaVenta) CrearPersona() IPersona {
	return &Usuario{}
}
