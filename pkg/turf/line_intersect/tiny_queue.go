package line_intersect

type CompareFunc func(a, b any) int

type TinyQueue struct {
	data    []any
	length  int
	compare CompareFunc
}

func NewTinyQueue(data []any, compare CompareFunc) *TinyQueue {
	if compare == nil {
		compare = defaultCompare
	}
	q := &TinyQueue{
		data:    data,
		length:  len(data),
		compare: compare,
	}
	if q.length > 0 {
		for i := (q.length >> 1) - 1; i >= 0; i-- {
			q.down(i)
		}
	}
	return q
}

func (q *TinyQueue) Push(item any) {
	q.data = append(q.data, item)
	q.length++
	q.up(q.length - 1)
}

func (q *TinyQueue) Pop() any {
	if q.length == 0 {
		return nil
	}
	top := q.data[0]
	bottom := q.data[q.length-1]
	q.data = q.data[:q.length-1]
	q.length--

	if q.length > 0 {
		q.data[0] = bottom
		q.down(0)
	}
	return top
}

func (q *TinyQueue) Peek() any {
	if q.length == 0 {
		return nil
	}
	return q.data[0]
}

func (q *TinyQueue) up(pos int) {
	item := q.data[pos]
	for pos > 0 {
		parent := (pos - 1) >> 1
		current := q.data[parent]
		if q.compare(item, current) >= 0 {
			break
		}
		q.data[pos] = current
		pos = parent
	}
	q.data[pos] = item
}

func (q *TinyQueue) down(pos int) {
	item := q.data[pos]
	halfLength := q.length >> 1

	for pos < halfLength {
		left := (pos << 1) + 1
		right := left + 1
		best := left

		if right < q.length && q.compare(q.data[right], q.data[left]) < 0 {
			best = right
		}
		if q.compare(q.data[best], item) >= 0 {
			break
		}

		q.data[pos] = q.data[best]
		pos = best
	}
	q.data[pos] = item
}

func defaultCompare(a, b any) int {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			return a - b
		}
	case float64:
		if b, ok := b.(float64); ok {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		}
	case string:
		if b, ok := b.(string); ok {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		}
	}
	return 0
}
