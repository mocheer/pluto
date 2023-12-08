package ds_arcgis_json

type Ring [][]float64
type Polygon []Ring

type Geometry struct {
	Points [][]float64   `json:"points"`
	Paths  [][][]float64 `json:"paths"`
	Rings  []Ring        `json:"rings"`
}

type SpatialReference struct {
	WKID       int
	LatestWKID int
}

type ArcGISFeature struct {
	Attributes map[string]any `json:"attributes"`
	Geometry   Geometry       `json:"geometry"`
	X          float64        `json:"x"`
	Y          float64        `json:"y"`
	Z          float64        `json:"z"`
	Xmin       float64        `json:"xmin"`
	Xmax       float64        `json:"xmax"`
	Ymin       float64        `json:"ymin"`
	Ymax       float64        `json:"ymax"`
	Paths      [][][]float64  `json:"paths"`
	Points     [][]float64    `json:"points"`
	Rings      []Ring         `json:"rings"`
}

type ArcGISJSON struct {
	DisplayFieldName string            `json:"displayFieldName"`
	FieldAliases     map[string]string `json:"fieldAliases"`
	GeometryType     string
	SpatialReference SpatialReference `json:"spatialReference"`
	Fields           []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Alias  string `json:"alias"`
		Length int    `json:"length"`
	} `json:"fields"`
	Features []ArcGISFeature `json:"features"`
}
