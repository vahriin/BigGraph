package types

type Edge struct {
	Nodes []Node
}

type Area struct {
	Border Bounds
	Edges []Edge
}

