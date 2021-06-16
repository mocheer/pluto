package tileset

import (
	"encoding/json"
)

const (
	TILE_REFINE_ADD     = "ADD"
	TILE_REFINE_REPLACE = "REPLACE"
)

var (
	TileDefaultTransform = [16]float64{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}
)

type Tileset struct {
	Asset              Asset              `json:"asset"`
	GeometricError     float64            `json:"geometricError"`
	Root               Tile               `json:"root"`
	Properties         *map[string]Schema `json:"properties,omitempty"`
	ExtensionsUsed     *[]string          `json:"extensionsUsed,omitempty"`
	ExtensionsRequired *[]string          `json:"extensionsRequired,omitempty"`
}

// ToJSON
func (ts *Tileset) ToJSON() (string, error) {
	b, e := json.Marshal(ts)
	return string(b), e
}
