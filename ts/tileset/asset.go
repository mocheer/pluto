package tileset

type Asset struct {
	Version        string `json:"version"`
	TilesetVersion string `json:"tilesetVersion,omitempty"`
	GltfUpAxis     string `json:"gltfUpAxis,omitempty"`
}
