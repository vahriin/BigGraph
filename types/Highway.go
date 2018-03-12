package types

// Highway is a road for car driving
type Highway struct {
	NodesID []uint64
}

/*func (h Highway) incidentNodes(nodeIndex int) []uint64 {
    incNodes := make([]uint64, 4)
    if i - 1 >= 0 {
        incNodes = append(incNodes, )
    } 

    if i + 1 < len(way.Refs) {
        al.AL[nd.Ref] = append(al.AL[nd.Ref], way.Refs[i + 1].Ref)
    }
}*/