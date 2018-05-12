package factorymethod

type Isoceles struct {
	Base triangulo
}

func (t *Isoceles) Nombrar() string {
	return "Soy un triangulo Isoceles"
}
