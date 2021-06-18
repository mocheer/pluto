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

// GetNext
func (m *Ring) GetNext() interface{} {
	data := m.data
	len := len(data)
	m.index++
	if m.index >= len {
		m.index = 0
	}
	return data[m.index]
}

//GetCurrent
func (m *Ring) GetCurrent() interface{} {
	data := m.data
	return data[m.index]
}

//SetIndex
func (m *Ring) SetIndex(index int) {
	m.index = index
}

//GetIndex
func (m *Ring) GetIndex() interface{} {
	return m.index
}
