package prototype

// Node a document object model node
type Node interface {
	String() string
	Parent() Node
	SetParent(node Node)
	Children() []Node
	AddChild(child Node)
	Clone() Node
}
