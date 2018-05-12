package factorymethod

type Escaleno struct {
	Base triangulo
}

func (t *Escaleno) Nombrar() string {
	return "Soy un triangulo Escaleno"
}
