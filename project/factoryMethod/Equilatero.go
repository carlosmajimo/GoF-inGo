package factorymethod

type Equilatero struct {
	Base triangulo
}

func (t *Equilatero) Nombrar() string {
	return "Soy un triangulo Equilatero"
}
