package adapter

type AcousticGuitar struct{}

func (a *AcousticGuitar) LeaveGuitar() string {
	return "Estoy cansado de tocar la guitarra acustica"
}

func (a *AcousticGuitar) PlayGuitar() string {
	return "Voy a tocar la guitarra acustica"
}

type AcousticGuitarAdapter struct {
	Acustica AcousticGuitar
}

func (a *AcousticGuitarAdapter) OnGuitar() string {
	return a.Acustica.PlayGuitar()
}

func (a *AcousticGuitarAdapter) OffGuitar() string {
	return a.Acustica.LeaveGuitar()
}
