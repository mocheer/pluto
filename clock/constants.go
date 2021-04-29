package clock

// Minute
// 1min=60s
var Minute = 60

// Hour
// 1h=60min=3600s
var Hour = 60 * Minute

// Day
// 1day=24h=86400s
var Day = 24 * Hour

// Month
// 1Mon=30Day
var Month = 30 * Day

// FmtDate 格式化日期
var FmtDate = "2006-01-02"

// FmtFullDate 格式化时间全量
var FmtFullDate = "2006-01-02 15:04:05"

// FmtCompactDate 格式化紧凑型日期
var FmtCompactDate = "20060102"
