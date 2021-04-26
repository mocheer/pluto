package clock

import "time"

// Parse
func Parse(str string, fmtStr string) (*Clock, error) {
	date, err := ParseTime(str, fmtStr)
	if err != nil {
		return nil, err
	}
	return &Clock{date}, nil
}

// Parse 解析时间字符串
func ParseTime(str string, fmtStr string) (time.Time, error) {
	return time.Parse(fmtStr, str)
}

// MustParse 解析时间字符串
func MustParse(str string, fmtStr string) *Clock {
	c, err := Parse(str, fmtStr)
	if err != nil {
		panic(`解析时间字符串失败!`)
	}
	return c
}
