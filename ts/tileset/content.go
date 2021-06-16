package tileset

// Content 不同版本 url(<1.0) 和 uri(1.0)
// url 一般是相对的，相对于引用的相对于引用的tileset文件
type Content struct {
	Url            string          `json:"url"`
	Uri            string          `json:"uri"`
	BoundingVolume *BoundingVolume `json:"boundingVolume,omitempty"`
}
