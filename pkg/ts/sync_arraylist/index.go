package sync_arraylist

import (
	"sync"

	"github.com/emirpasic/gods/lists/arraylist"
)

type SyncArrayList struct {
	sync.Mutex
	*arraylist.List
}

func New() *SyncArrayList {
	list := arraylist.New()
	return &SyncArrayList{List: list}
}

func (m *SyncArrayList) Add(values ...any) *SyncArrayList {
	m.Lock()
	defer m.Unlock()
	m.List.Add(values...)
	return m
}

func (m *SyncArrayList) TakeOut() any {
	m.Lock()
	defer m.Unlock()
	val, flag := m.List.Get(0)
	if flag {
		m.List.Remove(0)
	}
	return val
}

func (m *SyncArrayList) Takeouts() []any {
	m.Lock()
	defer m.Unlock()
	if m.List.Size() > 0 {
		vals := m.List.Values()
		m.List.Clear()
		return vals
	}
	return nil
}

func (m *SyncArrayList) EachTakeout(fn func(index int, value any)) *SyncArrayList {
	m.Lock()
	defer m.Unlock()
	//
	m.List.Each(fn)
	m.List.Clear()
	return m
}

// Size 容量
func (m *SyncArrayList) Size() int {
	m.Lock()
	defer m.Unlock()
	return m.List.Size()
}
