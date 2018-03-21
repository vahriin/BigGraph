package types

import "github.com/vahriin/BigGraph/lib/coordinates"

// Meta is type contains all required information from parsed file
type Meta struct {
	Bounds Bounds `xml:"bounds"`
	Nodes  []Node `xml:"node"`
	Ways   []Way  `xml:"way"`
}

// interpolationSearch use the Interpolation search
func (meta Meta) interpolationSearch(id uint64) Node {
	var left uint64
	var right = uint64(len(meta.Nodes)) - 1

	for meta.Nodes[left].ID < id && id < meta.Nodes[right].ID {
		mid := left + (id-meta.Nodes[left].ID)*(right-left)/(meta.Nodes[right].ID-meta.Nodes[left].ID)
		if meta.Nodes[mid].ID < id {
			left = mid + 1
		} else if meta.Nodes[mid].ID > id {
			right = mid - 1
		} else {
			return meta.Nodes[mid]
		}
	}

	if meta.Nodes[left].ID == id {
		return meta.Nodes[left]
	} else if meta.Nodes[right].ID == id {
		return meta.Nodes[right]
	} else {
		return Node{}
	}
}

// AdjList build the highways graph from Meta data
func (meta Meta) AdjList() AdjList {
	rectMin := meta.Bounds.Mins()
	rectMax := meta.Bounds.Maxs()

	var al AdjList
	al.AL = make(map[uint64][]uint64)
	al.Nodes = make(map[uint64]coordinates.GeneralCoords)

	for _, way := range meta.Ways {
		if way.IsHighway() {
			for i, nd := range way.Refs {
				if _, ok := al.AL[nd.Ref]; !ok {
					al.AL[nd.Ref] = way.IncidentNodes(i)

					//TODO:
					node := meta.interpolationSearch(nd.Ref)
					nodeEC := node.EuclidCoords()

					nodeEC.X -= rectMin.X
					nodeEC.Y = rectMax.Y - nodeEC.Y

					al.Nodes[nd.Ref] = coordinates.GeneralCoords{Earth: node.GeographicCoords, Euclid: nodeEC}
				} else {
					al.AL[nd.Ref] = append(al.AL[nd.Ref], way.IncidentNodes(i)...)
				}
			}
		}
	}

	return al
}
