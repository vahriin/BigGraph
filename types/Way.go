package types

type Way struct {
	Refs []Nd  `xml:"nd"`
	Tags []Tag `xml:"tag"`
}

func (way Way) IsHighway() bool {
	for _, tag := range way.Tags {
		if tag.Key == "highway" {
			if tag.Value == "footway" || tag.Value == "cycleway" || tag.Value == "bridleway" ||
				tag.Value == "living_street" || tag.Value == "pedestrian" || tag.Value == "steps" {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func (way Way) NodesId() []uint64 {
	nodesId := make([]uint64, 0, 10)

	for _, ref := range way.Refs {
		nodesId = append(nodesId, ref.Ref)
	}
	return nodesId
}
