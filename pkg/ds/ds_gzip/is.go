package ds_gzip

// IsGzip
// gzip文件头，前两个字符分别是1f 8b
func IsGzip(data []byte) bool {
	if data[0] != 0x1f || data[1] != 0x8b {
		return false
	}
	return true
}
