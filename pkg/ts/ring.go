package ts

// Ring
type Ring struct {
	data  []interface{}
	index int
}

// NewRing
func NewRing(data []interface{}) *Ring {
	ra := &Ring{data, -1}
	return ra
}

// Next
func (m *Ring) Next() interface{} {
	data := m.data
	len := len(data)
	m.index++
	if m.index >= len {
		m.index = 0
	}
	return data[m.index]
}

//Current
func (m *Ring) Current() interface{} {
	return m.data[m.index]
}

//Index
func (m *Ring) Index() interface{} {
	return m.index
}

//SetIndex
func (m *Ring) SetIndex(index int) {
	m.index = index
}
