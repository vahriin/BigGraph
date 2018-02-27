package xmlparse

type Edge struct {
	Nodes []Node
}

type Area struct {
	Border Bounds
	Edges []Edge
}

