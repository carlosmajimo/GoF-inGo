package strategy

type ITipoConduccion interface {
	ObtenerDescripcion() string
	ObtenerPotencia(decilitrosCombustible float64) int
	ObtenerIncrementoVelocidad(decilitrosCombustible float64) int
}
