package bridge

// Drawer draws on the underlying graphics device
type Drawer interface {
	// DrawEllipseInRect draws an ellipse in rectanlge
	DrawEllipseInRect(r Rect) string
}
