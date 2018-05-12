package bridge

import "fmt"

// OpenGL drawer
type OpenGL struct{}

// DrawEllipseInRect dibuja el elipse en el rectangulo
func (gl *OpenGL) DrawEllipseInRect(r Rect) string {
	return fmt.Sprintf("OpenGL está dibujando el elipse en el rectangulo %v", r)
}
