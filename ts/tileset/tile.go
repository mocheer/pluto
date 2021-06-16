package tileset

type Tile struct {
	Content             *Content        `json:"content,omitempty"`
	BoundingVolume      BoundingVolume  `json:"boundingVolume,omitempty"`
	ViewerRequestVolume *BoundingVolume `json:"viewerRequestVolume,omitempty"`
	GeometricError      float64         `json:"geometricError"`
	Refine              string          `json:"refine"`
	Transform           *[16]float64    `json:"transform,omitempty"`
	Children            []Tile          `json:"children,omitempty"`
}
