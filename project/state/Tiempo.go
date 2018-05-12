package state

type Tiempo struct {
	descripcion string
	Contexto    ITipoTiempo
}

func (t *Tiempo) SetContexto(tipo ITipoTiempo) {
	t.Contexto = tipo
}

func (t *Tiempo) Request() string{
	return t.Contexto.EjecutarAccion()
}
