package clock

import (
	"time"
)

// Clock 结构体
type Clock struct {
	date time.Time
}

// New 实例化
func New(date time.Time) *Clock {
	return &Clock{date}
}

// Now 实例化
func Now() *Clock {
	return New(time.Now())
}

// SinceHours 获取距离当前时间的小时数
func (c *Clock) SinceHours() float64 {
	return time.Since(c.date).Hours()
}

// SinceDays 获取距离当前时间的天数
func (c *Clock) SinceDays() float64 {
	return c.SinceHours() / 24
}

// SinceMonths 获取距离当前时间的月数
func (c *Clock) SinceMonths() float64 {
	return c.SinceDays() / 30
}

// Fmt 格式化返回时间字符串
func (c *Clock) Fmt(layout string) string {
	if layout == "" {
		layout = FmtDate
	}
	return c.date.Format(layout)
}

// Val 获取时间戳
func (c *Clock) Val() int64 {
	return c.date.Unix()
}
