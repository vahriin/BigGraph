package types

type Way struct {
	Refs []Nd  `xml:"nd"`
	Tags []Tag `xml:"tag"`
}

func (way Way) IsHighway() bool {
	for _, tag := range way.Tags {
		if tag.Key == "highway" {
			if tag.Value == "footway" || tag.Value == "cycleway" || tag.Value == "bridleway" ||
				tag.Value == "living_street" || tag.Value == "pedestrian" || tag.Value == "steps" ||
				tag.Value == "path" {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func (way Way) Edge() Edge {
	var edge Edge
	edge.NodesId = make([]uint64, 0, len(way.Refs))

	for _, ref := range way.Refs {
		edge.NodesId = append(edge.NodesId, ref.Ref)
	}
	return edge
}
