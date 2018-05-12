package state

import "fmt"

type Lluvioso struct{}

func (l *Lluvioso) EjecutarAccion() string{
	return fmt.Sprint("Está lloviendo")
}

type Nublado struct{}

func (l *Nublado) EjecutarAccion() string{
	return fmt.Sprint("Es un dia nublado")
}

type Nevado struct{}

func (l *Nevado) EjecutarAccion() string{
	return fmt.Sprint("Está Nevando")
}

type Soleado struct{}

func (l *Soleado) EjecutarAccion() string{
	return fmt.Sprint("Es un dia soleado")
}
