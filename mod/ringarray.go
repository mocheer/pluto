package model

// RingArray implements the RingArray interface.
type RingArray struct {
	data  []interface{}
	index int
}

// NewRingArray creates a new ring array.
func NewRingArray(data []interface{}) *RingArray {
	ra := &RingArray{data, -1}
	return ra
}

// GetNext
func (m *RingArray) GetNext() interface{} {
	data := m.data
	len := len(data)
	m.index++
	if m.index >= len {
		m.index = 0
	}
	return data[m.index]
}

//GetCurrent
func (m *RingArray) GetCurrent() interface{} {
	data := m.data
	return data[m.index]
}

//SetIndex
func (m *RingArray) SetIndex(index int) {
	m.index = index
}

//GetIndex
func (m *RingArray) GetIndex() interface{} {
	return m.index
}
