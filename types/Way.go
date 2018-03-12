package types

// Way is data from "way" tag
type Way struct {
	Refs []Nd  `xml:"nd"`
	Tags []Tag `xml:"tag"`
}

// IsHighway return true, if this highway can be used for driving a car
func (way Way) IsHighway() bool {
	for _, tag := range way.Tags {
		if tag.Key == "highway" {
			if !(tag.Value == "footway" || tag.Value == "cycleway" || tag.Value == "bridleway" ||
				tag.Value == "living_street" || tag.Value == "pedestrian" || tag.Value == "steps") {
				return true
			}
		}
	}
	return false
}

// IncidentNodes returns incident nodes of node with nodeIndex in Refs array
func (way Way) IncidentNodes(nodeIndex int) []uint64 {
	incNodes := make([]uint64, 0, 4)

	if nodeIndex-1 >= 0 {
		incNodes = append(incNodes, way.Refs[nodeIndex-1].Ref)
	}

	if nodeIndex+1 < len(way.Refs) {
		incNodes = append(incNodes, way.Refs[nodeIndex+1].Ref)
	}
	return incNodes
}

// Edge return array of Node's Id of this way
func (way Way) Edge() Highway {
	var edge Highway
	edge.NodesID = make([]uint64, 0, len(way.Refs))

	for _, ref := range way.Refs {
		edge.NodesID = append(edge.NodesID, ref.Ref)
	}
	return edge
}
