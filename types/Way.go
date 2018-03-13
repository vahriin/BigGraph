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
			if tag.Value == "motorway" || tag.Value == "motorway_link" ||
				tag.Value == "trunk" || tag.Value == "trunk_link" ||
				tag.Value == "primary" || tag.Value == "primary_link" ||
				tag.Value == "secondary" || tag.Value == "secondary_link" ||
				tag.Value == "tertiary" || tag.Value == "tertiary_link" ||
				tag.Value == "unclassified" ||
				tag.Value == "road" ||
				//tag.Value == "service" ||
				//tag.Value == "living_street" ||
				tag.Value == "residential" {
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
