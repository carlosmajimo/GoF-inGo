package bridge

type Circulo struct {
	DrawingContext Drawer
	Center         Point
	Radius         float64
}

// Draw draws a circle
func (circle *Circulo) Draw() string {
	rect := Rect{
		Location: Point{
			X: circle.Center.X - circle.Radius,
			Y: circle.Center.Y - circle.Radius,
		},
		Size: Size{
			Width:  2 * circle.Radius,
			Height: 2 * circle.Radius,
		},
	}

	return circle.DrawingContext.DrawEllipseInRect(rect)
}
