package clock

import "time"

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

// SinceLastHours 获取时间戳
func (c *Clock) SinceLastHours(num float64) bool {
	return time.Since(c.date).Hours()/num > 1
}

// SinceLastDays 获取时间戳
func (c *Clock) SinceLastDays(num float64) bool {
	return c.SinceLastHours(24 * num)
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
