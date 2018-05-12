package abstractFactory

type IInventario interface {
	Crear() string
}

type Inventario struct {
	Codigo   string
	Nombre   string
	Cantidad int
}

type Suministro struct {
	Base Inventario
	Tipo string
}

type Producto struct {
	Base Inventario
	Lote string
}

func (i *Suministro) Crear() string {
	return "Soy un suministro listo para usarse"
}

func (i *Producto) Crear() string {
	return "Soy un Producto listo para venderse"
}
