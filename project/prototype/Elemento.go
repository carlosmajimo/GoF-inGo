package prototype

import (
	"bytes"
	"fmt"
)

// Elemento representa la estructura de un elemento
type Elemento struct {
	texto    string
	parent   Node
	children []Node
}

// NewElemento crea nuevos elementos
func NewElemento(texto string) *Elemento {
	return &Elemento{
		texto:    texto,
		parent:   nil,
		children: make([]Node, 0),
	}
}

// NewDirector crea nuevos elementos
func NewDirector() *Elemento {
	return &Elemento{
		texto:    "Director de ingenieria",
		parent:   nil,
		children: make([]Node, 0),
	}
}

// NewManager crea nuevos elementos
func NewManager() *Elemento {
	return &Elemento{
		texto:    "Manager de ingenieria",
		parent:   nil,
		children: make([]Node, 0),
	}
}

// NewIngeniero crea nuevos elementos
func NewIngeniero() *Elemento {
	return &Elemento{
		texto:    "Ingeniero de software",
		parent:   nil,
		children: make([]Node, 0),
	}
}

// NewLiderOficina crea nuevos elementos
func NewLiderOficina() *Elemento {
	return &Elemento{
		texto:    "Lider de oficina",
		parent:   nil,
		children: make([]Node, 0),
	}
}

// Parent retorna el padre del elemento
func (e *Elemento) Parent() Node {
	return e.parent
}

// SetParent cambia el padre del Elemento
func (e *Elemento) SetParent(node Node) {
	e.parent = node
}

// Children retorna el hijo del elemento
func (e *Elemento) Children() []Node {
	return e.children
}

// AddChild agrega un hijo al elemento
func (e *Elemento) AddChild(child Node) {
	copy := child.Clone()
	copy.SetParent(e)
	e.children = append(e.children, copy)
}

// Clone crea una copia del elemento.
func (e *Elemento) Clone() Node {
	copy := &Elemento{
		texto:    e.texto,
		parent:   nil,
		children: make([]Node, 0),
	}
	for _, child := range e.children {
		copy.AddChild(child)
	}
	return copy
}

// String retorna la representacion en cadena de texto del elemento
func (e *Elemento) String() string {
	buffer := bytes.NewBufferString(e.texto)

	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}
