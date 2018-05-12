package builder

import "encoding/json"

// JSON Message Builder is concrete builder
type JSONMessageBuilder struct {
	messageRecipient string
	messageText      string
}

func (b *JSONMessageBuilder) SetRecipiente(recipiente string) {
	b.messageRecipient = recipiente
}

func (b *JSONMessageBuilder) SetTexto(texto string) {
	b.messageText = texto
}

func (b *JSONMessageBuilder) Mensaje() (*Mensaje, error) {
	m := make(map[string]string)
	m["recipiente"] = b.messageRecipient
	m["mensaje"] = b.messageText

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return &Mensaje{Cuerpo: data, Formato: "JSON"}, nil
}
