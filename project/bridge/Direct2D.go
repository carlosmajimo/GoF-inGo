package bridge

import "fmt"

// Direct2D drawer
type Direct2D struct{}

// DrawEllipseInRect dibuja el elipse en el rectangulo
func (d2d *Direct2D) DrawEllipseInRect(r Rect) string {
	return fmt.Sprintf("Direct2D está dibujando el elipse en el rectangulo %v", r)
}
