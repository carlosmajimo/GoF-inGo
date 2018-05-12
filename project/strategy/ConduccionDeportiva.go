package strategy

type ConduccionDeportiva struct{}

func (c *ConduccionDeportiva) ObtenerDescripcion() string {
	return "Conduccion Deportiva"
}

func (c *ConduccionDeportiva) ObtenerPotencia(decilitrosCombustible float64) int {
	return int(decilitrosCombustible*0.987) + 5
}

func (c *ConduccionDeportiva) ObtenerIncrementoVelocidad(decilitrosCombustible float64) int {
	return int(decilitrosCombustible*0.618) + 3
}
