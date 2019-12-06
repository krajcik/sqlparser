package sqlparser

type BindLocation struct {
	Offset, Length int
}

func (pq *ParsedQuery) GetBindLocation() []BindLocation {
	if len(pq.bindLocations) == 0 {
		return []BindLocation{}
	}

	locs := make([]BindLocation, len(pq.bindLocations))
	for _, loc := range pq.bindLocations {
		locs = append(locs, BindLocation{loc.offset, loc.length})
	}

	return locs
}

func (pq *ParsedQuery) SetBindLocation(offset, length int) {
	if len(pq.bindLocations) == 0 || len(pq.bindLocations) == 1 {
		pq.bindLocations = []bindLocation{{offset: offset, length: length}}
	}
}
