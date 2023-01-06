package clock

// clock的最小单位为秒

// Minute
// 1min=60s
var Minute = 60

// Hour
// 1h=60min=3600s
var Hour = 60 * Minute

// Day
// 1day=24h=86400s
var Day = 24 * Hour

// Week
var Week = 7 * Day

// Month
// 1Mon=30Day
var Month = 30 * Day

// Year 一年的所有秒数
var Year = 365 * Day

// FmtDate 格式化日期
var FmtDate = "2006-01-02"

// FmtFullDate 格式化时间（全量）相当于YYYY-MM:DD HH:mm:ss
var FmtFullDate = "2006-01-02 15:04:05"

// FmtCompactDate 格式化紧凑型日期
var FmtCompactDate = "20060102"

// FmtCompactDate 格式化紧凑型日期(全量)
var FmtCompactFullDate = "20060102150405"
