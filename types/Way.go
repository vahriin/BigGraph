package types


type Way struct {
	Refs []Nd `xml:"nd"`
	Tags []Tag `xml:"tag"`
}

func (way Way) IsHighway() bool {
	for _, tag := range way.Tags {
		if tag.Key == "highway" {
			return true
		}
	}
	return false
}

func (way Way) NodesId() []uint {
	nodesId := make([]uint, 0, 10)

	for _, ref := range way.Refs {
		nodesId = append(nodesId, ref.Ref)
	}
	return nodesId
}