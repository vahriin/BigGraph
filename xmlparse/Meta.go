package xmlparse

type Meta struct {
	Bounds Bounds `xml:"bounds"`
	Nodes []Node `xml:"node"`
	Ways []Way `xml:"way"`
}

func (meta Meta) search(id uint) (Node, bool) {
	var left uint = 0
	var right = uint(len(meta.Nodes)) - 1

	for meta.Nodes[left].Id < id && id < meta.Nodes[right].Id {
		mid := left + (id - meta.Nodes[left].Id) * (right - left) / (meta.Nodes[right].Id - meta.Nodes[left].Id)
		if meta.Nodes[mid].Id < id {
			left = mid + 1
		} else if meta.Nodes[mid].Id > id {
			right = mid - 1
		} else {
			return meta.Nodes[mid], true
		}
	}

	if meta.Nodes[left].Id == id {
		return meta.Nodes[left], true
	} else if meta.Nodes[right].Id == id {
		return meta.Nodes[right], true
	} else {
		return Node{}, false
	}
}

func (meta Meta) Graph() Area {
	var area Area
	area.Border = meta.Bounds

	for _, way := range meta.Ways {
		if way.IsHighway() {
			var edge Edge

			nodesId := way.NodesId()

			edge.Nodes = make([]Node, 0, len(nodesId))
			for _, id := range nodesId {
				if node, ok := meta.search(id); ok {
					edge.Nodes = append(edge.Nodes, node)
				}
			}

			area.Edges = append(area.Edges, edge)
		}
	}
	return area
}