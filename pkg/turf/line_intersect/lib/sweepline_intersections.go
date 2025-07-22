package lib

import (
	"container/heap"
	"math"
)

// TinyQueue 实现基于堆的优先队列
type TinyQueue struct {
	data    []interface{}
	compare func(a, b interface{}) int
}

func NewTinyQueue(compare func(a, b interface{}) int) *TinyQueue {
	return &TinyQueue{
		data:    make([]interface{}, 0),
		compare: compare,
	}
}

func (tq *TinyQueue) Len() int { return len(tq.data) }
func (tq *TinyQueue) Less(i, j int) bool {
	return tq.compare(tq.data[i], tq.data[j]) < 0
}
func (tq *TinyQueue) Swap(i, j int) {
	tq.data[i], tq.data[j] = tq.data[j], tq.data[i]
}

func (tq *TinyQueue) Push(x interface{}) {
	tq.data = append(tq.data, x)
}

func (tq *TinyQueue) Pop() interface{} {
	n := len(tq.data)
	item := tq.data[n-1]
	tq.data = tq.data[0 : n-1]
	return item
}

func (tq *TinyQueue) Peek() interface{} {
	if len(tq.data) > 0 {
		return tq.data[0]
	}
	return nil
}

// 事件结构定义
type Event struct {
	X, Y           float64
	FeatureID      int
	RingID         int
	EventID        int
	OtherEvent     *Event
	IsLeftEndpoint bool
}

func (e *Event) IsSamePoint(other *Event) bool {
	return e.X == other.X && e.Y == other.Y
}

func (e *Event) AsNewXY() []float64 {
	return []float64{e.X, e.Y}
}

// 线段结构定义
type Segment struct {
	LeftSweepEvent  *Event
	RightSweepEvent *Event
}

// 常量定义
const (
	epsilon        = 1.1102230246251565e-16
	splitter       = 134217729
	resulterrbound = (3 + 8*epsilon) * epsilon
	ccwerrboundA   = (3 + 16*epsilon) * epsilon
	ccwerrboundB   = (2 + 12*epsilon) * epsilon
	ccwerrboundC   = (9 + 64*epsilon) * epsilon * epsilon
)

// 比较函数
func defaultCompare(a, b interface{}) int {
	f1 := a.(float64)
	f2 := b.(float64)
	if f1 < f2 {
		return -1
	}
	if f1 > f2 {
		return 1
	}
	return 0
}

func checkWhichEventIsLeft(e1, e2 interface{}) int {
	a := e1.(*Event)
	b := e2.(*Event)

	if a.X > b.X {
		return 1
	}
	if a.X < b.X {
		return -1
	}

	if a.X == b.X && (a.FeatureID != b.FeatureID || a.RingID != b.RingID) {
		if a.IsLeftEndpoint && !b.IsLeftEndpoint {
			return -1
		}
	}

	if a.Y != b.Y {
		if a.Y > b.Y {
			return 1
		}
		return -1
	}
	return 1
}

func checkWhichSegmentHasRightEndpointFirst(seg1, seg2 interface{}) int {
	a := seg1.(*Segment)
	b := seg2.(*Segment)

	if a.RightSweepEvent.X > b.RightSweepEvent.X {
		return 1
	}
	if a.RightSweepEvent.X < b.RightSweepEvent.X {
		return -1
	}

	if a.RightSweepEvent.Y != b.RightSweepEvent.Y {
		if a.RightSweepEvent.Y < b.RightSweepEvent.Y {
			return 1
		}
		return -1
	}
	return 1
}

// 填充事件队列
func fillEventQueue(coords [][][][]float64, eventQueue *TinyQueue) {
	featureID := 0
	ringID := 0
	eventID := 0

	for i := 0; i < len(coords); i++ {
		for ii := 0; ii < len(coords[i]); ii++ {
			currentRing := coords[i][ii]
			ringID++

			for iii := 0; iii < len(currentRing)-1; iii++ {
				currentP := currentRing[iii]
				nextP := currentRing[iii+1]

				e1 := &Event{
					X:         currentP[0],
					Y:         currentP[1],
					FeatureID: featureID,
					RingID:    ringID,
					EventID:   eventID,
				}

				e2 := &Event{
					X:         nextP[0],
					Y:         nextP[1],
					FeatureID: featureID,
					RingID:    ringID,
					EventID:   eventID + 1,
				}

				e1.OtherEvent = e2
				e2.OtherEvent = e1

				if checkWhichEventIsLeft(e1, e2) > 0 {
					e2.IsLeftEndpoint = true
					e1.IsLeftEndpoint = false
				} else {
					e1.IsLeftEndpoint = true
					e2.IsLeftEndpoint = false
				}

				heap.Push(eventQueue, e1)
				heap.Push(eventQueue, e2)

				eventID += 2
			}
		}
		featureID++
	}
}

// 精确计算函数
func estimate(elen int, e []float64) float64 {
	Q := e[0]
	for i := 1; i < elen; i++ {
		Q += e[i]
	}
	return Q
}

func sum(elen int, e []float64, flen int, f []float64, h []float64) int {
	var Q, Qnew, hh, bvirt float64
	enow := e[0]
	fnow := f[0]
	eindex := 0
	findex := 0

	if (fnow > enow) == (fnow > -enow) {
		Q = enow
		eindex++
		if eindex < elen {
			enow = e[eindex]
		}
	} else {
		Q = fnow
		findex++
		if findex < flen {
			fnow = f[findex]
		}
	}

	hindex := 0
	for eindex < elen && findex < flen {
		if (fnow > enow) == (fnow > -enow) {
			Qnew = enow + Q
			hh = Q - (Qnew - enow)
			eindex++
			if eindex < elen {
				enow = e[eindex]
			}
		} else {
			Qnew = fnow + Q
			hh = Q - (Qnew - fnow)
			findex++
			if findex < flen {
				fnow = f[findex]
			}
		}
		Q = Qnew
		if hh != 0 {
			h[hindex] = hh
			hindex++
		}
	}

	for eindex < elen {
		Qnew = Q + enow
		bvirt = Qnew - Q
		hh = Q - (Qnew - bvirt) + (enow - bvirt)
		eindex++
		if eindex < elen {
			enow = e[eindex]
		}
		Q = Qnew
		if hh != 0 {
			h[hindex] = hh
			hindex++
		}
	}

	for findex < flen {
		Qnew = Q + fnow
		bvirt = Qnew - Q
		hh = Q - (Qnew - bvirt) + (fnow - bvirt)
		findex++
		if findex < flen {
			fnow = f[findex]
		}
		Q = Qnew
		if hh != 0 {
			h[hindex] = hh
			hindex++
		}
	}

	if Q != 0 || hindex == 0 {
		h[hindex] = Q
		hindex++
	}
	return hindex
}

func orient2dadapt(ax, ay, bx, by, cx, cy, detsum float64) float64 {
	var acx = ax - cx
	var bcx = bx - cx
	var acy = ay - cy
	var bcy = by - cy

	s1 := acx * bcy
	c := splitter * acx
	ahi := c - (c - acx)
	alo := acx - ahi
	c = splitter * bcy
	bhi := c - (c - bcy)
	blo := bcy - bhi
	s0 := alo*blo - (s1 - ahi*bhi - alo*bhi - ahi*blo)

	t1 := acy * bcx
	c = splitter * acy
	ahi = c - (c - acy)
	alo = acy - ahi
	c = splitter * bcx
	bhi = c - (c - bcx)
	blo = bcx - bhi
	t0 := alo*blo - (t1 - ahi*bhi - alo*bhi - ahi*blo)

	_i := s0 - t0
	bvirt := s0 - _i
	B := []float64{s0 - (_i + bvirt) + (bvirt - t0), 0, 0, 0}
	_j := s1 + _i
	bvirt = _j - s1
	_0 := s1 - (_j - bvirt) + (_i - bvirt)
	_i = _0 - t1
	bvirt = _0 - _i
	B[1] = _0 - (_i + bvirt) + (bvirt - t1)
	u3 := _j + _i
	bvirt = u3 - _j
	B[2] = _j - (u3 - bvirt) + (_i - bvirt)
	B[3] = u3

	det := estimate(4, B)
	errbound := ccwerrboundB * detsum
	if det >= errbound || -det >= errbound {
		return det
	}

	bvirt = ax - acx
	acxtail := ax - (acx + bvirt) + (bvirt - cx)
	bvirt = bx - bcx
	bcxtail := bx - (bcx + bvirt) + (bvirt - cx)
	bvirt = ay - acy
	acytail := ay - (acy + bvirt) + (bvirt - cy)
	bvirt = by - bcy
	bcytail := by - (bcy + bvirt) + (bvirt - cy)

	if acxtail == 0 && acytail == 0 && bcxtail == 0 && bcytail == 0 {
		return det
	}

	errbound = ccwerrboundC*detsum + resulterrbound*math.Abs(det)
	det += (acx*bcytail + bcy*acxtail) - (acy*bcxtail + bcx*acytail)
	if det >= errbound || -det >= errbound {
		return det
	}

	s1 = acxtail * bcy
	c = splitter * acxtail
	ahi = c - (c - acxtail)
	alo = acxtail - ahi
	c = splitter * bcy
	bhi = c - (c - bcy)
	blo = bcy - bhi
	s0 = alo*blo - (s1 - ahi*bhi - alo*bhi - ahi*blo)

	t1 = acytail * bcx
	c = splitter * acytail
	ahi = c - (c - acytail)
	alo = acytail - ahi
	c = splitter * bcx
	bhi = c - (c - bcx)
	blo = bcx - bhi
	t0 = alo*blo - (t1 - ahi*bhi - alo*bhi - ahi*blo)

	_i = s0 - t0
	bvirt = s0 - _i
	u := []float64{s0 - (_i + bvirt) + (bvirt - t0), 0, 0, 0}
	_j = s1 + _i
	bvirt = _j - s1
	_0 = s1 - (_j - bvirt) + (_i - bvirt)
	_i = _0 - t1
	bvirt = _0 - _i
	u[1] = _0 - (_i + bvirt) + (bvirt - t1)
	u3 = _j + _i
	bvirt = u3 - _j
	u[2] = _j - (u3 - bvirt) + (_i - bvirt)
	u[3] = u3

	C1 := make([]float64, 8)
	C1len := sum(4, B, 4, u, C1)

	s1 = acx * bcytail
	c = splitter * acx
	ahi = c - (c - acx)
	alo = acx - ahi
	c = splitter * bcytail
	bhi = c - (c - bcytail)
	blo = bcytail - bhi
	s0 = alo*blo - (s1 - ahi*bhi - alo*bhi - ahi*blo)

	t1 = acy * bcxtail
	c = splitter * acy
	ahi = c - (c - acy)
	alo = acy - ahi
	c = splitter * bcxtail
	bhi = c - (c - bcxtail)
	blo = bcxtail - bhi
	t0 = alo*blo - (t1 - ahi*bhi - alo*bhi - ahi*blo)

	_i = s0 - t0
	bvirt = s0 - _i
	u[0] = s0 - (_i + bvirt) + (bvirt - t0)
	_j = s1 + _i
	bvirt = _j - s1
	_0 = s1 - (_j - bvirt) + (_i - bvirt)
	_i = _0 - t1
	bvirt = _0 - _i
	u[1] = _0 - (_i + bvirt) + (bvirt - t1)
	u3 = _j + _i
	bvirt = u3 - _j
	u[2] = _j - (u3 - bvirt) + (_i - bvirt)
	u[3] = u3

	C2 := make([]float64, 12)
	C2len := sum(C1len, C1, 4, u, C2)

	s1 = acxtail * bcytail
	c = splitter * acxtail
	ahi = c - (c - acxtail)
	alo = acxtail - ahi
	c = splitter * bcytail
	bhi = c - (c - bcytail)
	blo = bcytail - bhi
	s0 = alo*blo - (s1 - ahi*bhi - alo*bhi - ahi*blo)

	t1 = acytail * bcxtail
	c = splitter * acytail
	ahi = c - (c - acytail)
	alo = acytail - ahi
	c = splitter * bcxtail
	bhi = c - (c - bcxtail)
	blo = bcxtail - bhi
	t0 = alo*blo - (t1 - ahi*bhi - alo*bhi - ahi*blo)

	_i = s0 - t0
	bvirt = s0 - _i
	u[0] = s0 - (_i + bvirt) + (bvirt - t0)
	_j = s1 + _i
	bvirt = _j - s1
	_0 = s1 - (_j - bvirt) + (_i - bvirt)
	_i = _0 - t1
	bvirt = _0 - _i
	u[1] = _0 - (_i + bvirt) + (bvirt - t1)
	u3 = _j + _i
	bvirt = u3 - _j
	u[2] = _j - (u3 - bvirt) + (_i - bvirt)
	u[3] = u3

	D := make([]float64, 16)
	Dlen := sum(C2len, C2, 4, u, D)

	return D[Dlen-1]
}

func orient2d(ax, ay, bx, by, cx, cy float64) float64 {
	detleft := (ay - cy) * (bx - cx)
	detright := (ax - cx) * (by - cy)
	det := detleft - detright

	if detleft == 0 || detright == 0 || (detleft > 0) != (detright > 0) {
		return det
	}

	detsum := math.Abs(detleft + detright)
	if math.Abs(det) >= ccwerrboundA*detsum {
		return det
	}

	return -orient2dadapt(ax, ay, bx, by, cx, cy, detsum)
}

func testSegmentIntersect(seg1, seg2 *Segment) interface{} {
	if seg1 == nil || seg2 == nil {
		return false
	}

	x1 := seg1.LeftSweepEvent.X
	y1 := seg1.LeftSweepEvent.Y
	x2 := seg1.RightSweepEvent.X
	y2 := seg1.RightSweepEvent.Y
	x3 := seg2.LeftSweepEvent.X
	y3 := seg2.LeftSweepEvent.Y
	x4 := seg2.RightSweepEvent.X
	y4 := seg2.RightSweepEvent.Y

	score1 := orient2d(x1, y1, x2, y2, x3, y3)
	score2 := orient2d(x1, y1, x2, y2, x4, y4)

	if (score1 > 0 && score2 > 0) || (score1 < 0 && score2 < 0) {
		return false
	}

	if seg1.LeftSweepEvent.RingID == seg2.LeftSweepEvent.RingID {
		if seg1.RightSweepEvent.IsSamePoint(seg2.LeftSweepEvent) ||
			seg1.RightSweepEvent.IsSamePoint(seg2.RightSweepEvent) ||
			seg1.LeftSweepEvent.IsSamePoint(seg2.LeftSweepEvent) ||
			seg1.LeftSweepEvent.IsSamePoint(seg2.RightSweepEvent) {
			return false
		}
	} else {
		if seg1.RightSweepEvent.IsSamePoint(seg2.LeftSweepEvent) {
			return seg2.LeftSweepEvent.AsNewXY()
		}
		if seg1.RightSweepEvent.IsSamePoint(seg2.RightSweepEvent) {
			return seg2.RightSweepEvent.AsNewXY()
		}
		if seg1.LeftSweepEvent.IsSamePoint(seg2.LeftSweepEvent) {
			return seg2.LeftSweepEvent.AsNewXY()
		}
		if seg1.LeftSweepEvent.IsSamePoint(seg2.RightSweepEvent) {
			return seg2.RightSweepEvent.AsNewXY()
		}
	}

	denom := ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
	if denom == 0 {
		return false
	}

	numeA := ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
	numeB := ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))

	uA := numeA / denom
	uB := numeB / denom

	if uA >= 0 && uA <= 1 && uB >= 0 && uB <= 1 {
		x := x1 + uA*(x2-x1)
		y := y1 + uA*(y2-y1)
		return []float64{x, y}
	}
	return false
}

func runCheck(eventQueue *TinyQueue, ignoreSelfIntersections bool) [][]float64 {
	intersectionPoints := make([][]float64, 0)
	outQueue := NewTinyQueue(checkWhichSegmentHasRightEndpointFirst)
	heap.Init(outQueue)

	for eventQueue.Len() > 0 {
		event := heap.Pop(eventQueue).(*Event)
		if event.IsLeftEndpoint {
			segment := &Segment{
				LeftSweepEvent:  event,
				RightSweepEvent: event.OtherEvent,
			}

			for _, item := range outQueue.data {
				otherSeg := item.(*Segment)
				if ignoreSelfIntersections {
					if otherSeg.LeftSweepEvent.FeatureID == event.FeatureID {
						continue
					}
				}
				intersection := testSegmentIntersect(segment, otherSeg)
				if pts, ok := intersection.([]float64); ok {
					intersectionPoints = append(intersectionPoints, pts)
				}
			}
			heap.Push(outQueue, segment)
		} else {
			heap.Pop(outQueue)
		}
	}
	return intersectionPoints
}

func SweeplineIntersections(polygon [][][][]float64, ignoreSelfIntersections bool) [][]float64 {
	eventQueue := NewTinyQueue(checkWhichEventIsLeft)
	heap.Init(eventQueue)
	fillEventQueue(polygon, eventQueue)
	return runCheck(eventQueue, ignoreSelfIntersections)
}
