package adapter

type ElectricGuitar struct{}

func (e *ElectricGuitar) OnGuitar() string {
	return "Voy a tocar la guitarra electrica"
}

func (e *ElectricGuitar) OffGuitar() string {
	return "Estoy cansado de tocar la guitarra electrica"
}
