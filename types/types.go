package types

type Collection struct {
	// Collection name
	Name string `json:"name"`
	// Collection metadata dictionary
	Metadata map[string]interface{} `json:"metadata"`
}

type Segment interface {
	// Segment name
}

type VectorReader interface {
	// Segment name
}

type SegmentManager interface {
	// Create a new segment
	CreateSegment(segment Segment) error
}
