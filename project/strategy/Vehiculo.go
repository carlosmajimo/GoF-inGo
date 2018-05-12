package strategy

import "fmt"

type Vehiculo struct {
	Contexto ITipoConduccion
}

func (v *Vehiculo) ConduccionNormal() {
	v.Contexto = &ConduccionNormal{}
}

func (v *Vehiculo) ConduccionDeportiva() {
	v.Contexto = &ConduccionDeportiva{}
}

func (v *Vehiculo) Acelerar(combustible float64) string {
	descripcion := v.Contexto.ObtenerDescripcion()
	incrementoVelocidad := v.Contexto.ObtenerIncrementoVelocidad(combustible)
	potencia := v.Contexto.ObtenerPotencia(combustible)

	return fmt.Sprintf("Tipo de conducción: %s, con la potencia de %d se aumento la velocidad en %d", descripcion, potencia, incrementoVelocidad)
	//fmt.Printf("Combustible inyectado: %.2f\n", combustible)
	//fmt.Printf("Potencia proporcionada: %d\n", potencia)
	//fmt.Printf("Incremento de velocidad: %d\n\n", incrementoVelocidad)

}
