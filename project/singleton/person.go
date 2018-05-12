package singleton

import "sync"

type person struct {
	nombre string
}

func (p *person) GetNombre() string {
	return p.nombre
}

func (p *person) SetNombre(nombre string) {
	p.nombre = nombre
}

var (
	instance *person
	once     sync.Once
)

// Person retorna la instancia de la persona
func Person() *person {
	once.Do(func() {
		instance = &person{
			nombre: "",
		}
	})
	return instance
}

func Sumar(a, b int) int {
	return a + b
}
