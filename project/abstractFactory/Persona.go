package abstractFactory

type IPersona interface {
	Crear() string
}

type Persona struct {
	Codigo   string
	Nombre   string
	Cantidad int
}

type Usuario struct {
	Base Persona
	Edad int
}

type Proveedor struct {
	Base     Persona
	Telefono int
}

func (i *Usuario) Crear() string {
	return "Soy un usuario listo para comprar"
}

func (i *Proveedor) Crear() string {
	return "Soy un Proveedor listo para vender"
}
