package builder

import "encoding/xml"

// XML Message Builder is concrete builder
type XMLMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *XMLMessageBuilder) SetRecipiente(recipiente string) {
	b.messageRecipient = recipiente
}

func (b *XMLMessageBuilder) SetTexto(texto string) {
	b.messageText = texto
}

func (b *XMLMessageBuilder) Mensaje() (*Mensaje, error) {
	type XMLMessage struct {
		Recipiente string `xml:"recipiente"`
		Texto      string `xml:"cuerpo"`
	}

	m := XMLMessage{
		Recipiente: b.messageRecipient,
		Texto:      b.messageText,
	}

	data, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Mensaje{Cuerpo: data, Formato: "XML"}, nil
}
