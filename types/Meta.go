package types

import "strconv"

type Meta struct {
	Bounds Bounds `xml:"bounds"`
	Nodes  []Node `xml:"node"`
	Ways   []Way  `xml:"way"`
}

func (meta Meta) search(id uint64) (Node, bool) {
	var left uint64 = 0
	var right = uint64(len(meta.Nodes)) - 1

	for meta.Nodes[left].Id < id && id < meta.Nodes[right].Id {
		mid := left + (id-meta.Nodes[left].Id)*(right-left)/(meta.Nodes[right].Id-meta.Nodes[left].Id)
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
	rectMin := meta.Bounds.Mins()
	rectMax := meta.Bounds.Maxs()

	var area Area
	area.Edges = make([]Edge, 0, 15000)
	area.Points = make(map[uint64]GeneralCoords)

	for _, way := range meta.Ways {
		if way.IsHighway() {
			edge := way.Edge()
			for _, id := range edge.NodesId {
				if _, ok := area.Points[id]; !ok {
					if node, ok1 := meta.search(id); ok1 {
						nodeEC := node.EuclidCoords()
						nodeEC.X -= rectMin.X
						nodeEC.Y = rectMax.Y - nodeEC.Y

						area.Points[id] = GeneralCoords{Earth: node.EarthCoords, Euclid: nodeEC}
					} else {
						panic("No nodes with id: " + strconv.FormatUint(id, 10))
					}
				}
			}
			area.Edges = append(area.Edges, edge)
		}
	}
	return area
}
