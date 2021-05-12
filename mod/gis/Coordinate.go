package gis

import (
	"fmt"
	"math"
)

type Coordinate struct {
	Row  float64 //行
	Col  float64 //列
	Zoom float64 //缩放级别
}

/**
 * 返回一个包含当前坐标的新的 coordinate 对象。即：向下取整。
 */
func (m *Coordinate) Container() *Coordinate {
	return &Coordinate{math.Floor(m.Row), math.Floor(m.Col), m.Zoom}
}

/**
 * 将当前坐标缩放至 destination 级别，返回缩放后的对象副本。该方法不修改原始对象。
 * @param	destination	缩放后的缩放级别
 * @return	缩放后对应的坐标（行、列、缩放级别）
 */
func (m *Coordinate) ZoomTo(destination float64) *Coordinate {
	return &Coordinate{m.Row * math.Pow(2, destination-m.Zoom), m.Col * math.Pow(2, destination-m.Zoom), destination}
}

/**
 * 对当前坐标缩放 distance 级，返回缩放后对象副本。该方法不修改原始对象。
 * @param	distance	要进行缩放的等级(正数为放大，负数为缩小)
 * @return	缩放后对应的坐标（行、列、缩放级别）
 */
func (m *Coordinate) ZoomBy(distance float64) *Coordinate {
	return &Coordinate{m.Row * math.Pow(2, distance), m.Col * math.Pow(2, distance), m.Zoom + distance}
}

/**
 * 当前坐标是否恰为某行（即不含小数）
 * @return	如果恰为某行，则返回true，否则返回false
 */
func (m *Coordinate) isRowEdge() bool {
	return math.Floor(m.Row) == m.Row
}

/**
 * 当前坐标是否恰为某列（即不含小数）
 * @return	如果恰为某列，则返回true，否则返回false
 */
func (m *Coordinate) isColumnEdge() bool {
	return math.Floor(m.Col) == m.Col
}

/**
 * 当前坐标是否恰为某行列（即不含小数）
 * @return	如果恰为某行某列，则返回true，否则返回false
 */
func (m *Coordinate) isEdge() bool {
	return m.isRowEdge() && m.isColumnEdge()
}

/**
 * 返回当前坐标上移 distance 行对应的坐标，即：Row - distance
 * @param	distance	往上的行数
 * @return			对应坐标
 */
func (m *Coordinate) up(distance float64) *Coordinate {

	return &Coordinate{m.Row - distance, m.Col, m.Zoom}
}

/**
 * 返回当前坐标右移 distance 列对应的坐标，即：Column + distance
 * @param	distance	往右的列数
 * @return			对应坐标
 */
func (m *Coordinate) right(distance float64) *Coordinate { // = 1
	return &Coordinate{m.Row, m.Col + distance, m.Zoom}
}

/**
 * 返回当前坐标下移 distance 行对应的坐标，即：Row + distance
 * @param	distance	下移的行数
 * @return			对应坐标
 */
func (m *Coordinate) down(distance float64) *Coordinate { //
	return &Coordinate{m.Row + distance, m.Col, m.Zoom}
}

/**
 * 返回当前坐标左移 distance 列对应的坐标，即：Column - distance
 * @param	distance	往左的列数
 * @return			对应坐标
 */
func (m *Coordinate) left(distance float64) *Coordinate {
	return &Coordinate{m.Row, m.Col - distance, m.Zoom}
}

/**
 * 如果两个坐标对象表示同一个位置，则返回true，否则返回false.
 */
func (m *Coordinate) equalTo(coord *Coordinate) bool {
	return coord != nil && coord.Row == m.Row && coord.Col == m.Col && coord.Zoom == m.Zoom
}

/**
 * 返回当前坐标对应的副本
 * @return	当前坐标对应的副本
 */
func (m *Coordinate) clone() *Coordinate {
	return &Coordinate{m.Row, m.Col, m.Zoom}
}

/**
 * 返回坐标的字符串表示形式
 * @return	坐标的字符串表示形式，格式为：Column,Row,Zoom
 */
func (m *Coordinate) String() string {
	return fmt.Sprintf("%d,%d,%d", m.Row, m.Col, m.Zoom)
}
