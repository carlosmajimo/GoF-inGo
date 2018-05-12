package decorator

import "fmt"

// Repositorio de datos
type Repositorio struct{}

func (r *Repositorio) Fetch(args Args) (Data, string) {
	if len(args) == 0 {
		return Data{}, fmt.Sprintf("No se han recibido argumentos")
	}

	data := Data{
		"user":     "root",
		"password": "swordfish",
	}
	//fmt.Printf("El repositorio a retornado los datos con exito: %v\n", data)
	return data, ""
}
