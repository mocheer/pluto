package ts

// Ring
type Ring struct {
	data  []any
	index int
}

// NewRing
func NewRing(data []any) *Ring {
	ra := &Ring{data, -1}
	return ra
}

// Next
func (m *Ring) Next() any {
	data := m.data
	len := len(data)
	m.index++
	if m.index >= len {
		m.index = 0
	}
	return data[m.index]
}

//Current
func (m *Ring) Current() any {
	return m.data[m.index]
}

//Index
func (m *Ring) Index() any {
	return m.index
}

//SetIndex
func (m *Ring) SetIndex(index int) {
	m.index = index
}
