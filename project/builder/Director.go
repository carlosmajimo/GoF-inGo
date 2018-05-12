package builder

type Sender struct{}

func (s *Sender) BuildMessage(builder MensajeBuilder) (*Mensaje, error) {
	builder.SetRecipiente("Santa Claus")
	builder.SetTexto("Estoy cansado de entregar regalos cada año")
	return builder.Mensaje()
}
