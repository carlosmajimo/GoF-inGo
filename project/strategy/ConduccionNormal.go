package strategy

type ConduccionNormal struct{}

// ObtenerDescripcion retorna la descripcion del tipo de conduccion
func (c *ConduccionNormal) ObtenerDescripcion() string {
	return "Conduccion Normal"
}

func (c *ConduccionNormal) ObtenerPotencia(decilitrosCombustible float64) int {
	return int(decilitrosCombustible*0.842) + 5
}

func (c *ConduccionNormal) ObtenerIncrementoVelocidad(decilitrosCombustible float64) int {
	return int(decilitrosCombustible*0.422) + 3
}
