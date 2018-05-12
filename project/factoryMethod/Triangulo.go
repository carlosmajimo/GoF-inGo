package factorymethod

type ITriangulo interface {
	Nombrar() string
}

type triangulo struct {
	L1, L2, L3 int
}
