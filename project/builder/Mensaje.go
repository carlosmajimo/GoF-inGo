package builder

type Mensaje struct {
	// Message Body
	Cuerpo []byte
	// Message Format
	Formato string
}

type MensajeBuilder interface {
	// Set the message's recipient
	SetRecipiente(recipient string)
	// Set the message's text
	SetTexto(texto string)
	// Returns the built Message
	Mensaje() (*Mensaje, error)
}
