package sTime

import "time"

// Timezone 获取时区
func (c STime) Timezone() string {
	return c.Location().String()
}

// DaysInYear 获取本年的总天数
func (c STime) DaysInYear() int {
	if c.IsZero() {
		return 0
	}
	days := DaysPerNormalYear
	if c.IsLeapYear() {
		days = DaysPerLeapYear
	}
	return days
}

// DaysInMonth 获取本月的总天数
func (c STime) DaysInMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.EndOfMonth().Day()
}

// MonthOfYear 获取本年的第几月(从1开始)
func (c STime) MonthOfYear() int {
	if c.IsZero() {
		return 0
	}
	return int(c.Time.Month())
}

// DayOfYear 获取本年的第几天(从1开始)
func (c STime) DayOfYear() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.YearDay()
}

// DayOfMonth 获取本月的第几天(从1开始)
func (c STime) DayOfMonth() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Day()
}

// DayOfWeek 获取本周的第几天(从1开始)
func (c STime) DayOfWeek() int {
	if c.IsZero() {
		return 0
	}
	day := int(c.Time.Weekday())
	if day == 0 {
		return DaysPerWeek
	}
	return day
}

// WeekOfYear 获取本年的第几周(从1开始)
func (c STime) WeekOfYear() int {
	if c.IsZero() {
		return 0
	}
	_, week := c.Time.ISOWeek()
	return week
}

// WeekOfMonth 获取本月的第几周(从1开始)
func (c STime) WeekOfMonth() int {
	if c.IsZero() {
		return 0
	}
	day := c.Time.Day()
	if day < DaysPerWeek {
		return 1
	}
	return day%DaysPerWeek + 1
}

// Year 获取当前年
func (c STime) Year() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Year()
}

// Quarter 获取当前季度
func (c STime) Quarter() int {
	if c.IsZero() {
		return 0
	}
	switch {
	case c.Month() >= 10:
		return 4
	case c.Month() >= 7:
		return 3
	case c.Month() >= 4:
		return 2
	case c.Month() >= 1:
		return 1
	default:
		return 0
	}
}

// Month 获取当前月
func (c STime) Month() int {
	if c.IsZero() {
		return 0
	}
	return c.MonthOfYear()
}

// Week 获取当前周(从0开始)
func (c STime) Week() int {
	if c.IsZero() {
		return -1
	}
	return int(c.Time.Weekday())
}

// Day 获取当前日
func (c STime) Day() int {
	if c.IsZero() {
		return 0
	}
	return c.DayOfMonth()
}

// Hour 获取当前小时
func (c STime) Hour() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Hour()
}

// Minute 获取当前分钟数
func (c STime) Minute() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Minute()
}

// Second 获取当前秒数
func (c STime) Second() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Second()
}

// Millisecond 获取当前毫秒数
func (c STime) Millisecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Nanosecond() / 1e6
}

// Microsecond 获取当前微秒数
func (c STime) Microsecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Nanosecond() / 1e9
}

// Nanosecond 获取当前纳秒数
func (c STime) Nanosecond() int {
	if c.IsZero() {
		return 0
	}
	return c.Time.Nanosecond()
}

// ------------------------------------------

// StartOfYear 本年开始时间
func (c STime) StartOfYear() time.Time {
	return time.Date(c.Time.Year(), 1, 1, 0, 0, 0, 0, c.Location())
}

// EndOfYear 本年结束时间
func (c STime) EndOfYear() time.Time {
	return time.Date(c.Time.Year(), 12, 31, 23, 59, 59, 0, c.Location())
}

// StartOfMonth 本月开始时间
func (c STime) StartOfMonth() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), 1, 0, 0, 0, 0, c.Location())
}

// EndOfMonth 本月结束时间
func (c STime) EndOfMonth() time.Time {
	t := time.Date(c.Time.Year(), c.Time.Month(), 1, 23, 59, 59, 0, c.Location())
	return t.AddDate(0, 1, -1)
}

// StartOfWeek 本周开始时间
func (c STime) StartOfWeek() time.Time {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}

	t := time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 0, 0, 0, 0, c.Location())
	return t.AddDate(0, 0, int(1-days))
}

// EndOfWeek 本周结束时间
func (c STime) EndOfWeek() time.Time {
	days := c.Time.Weekday()
	if days == 0 {
		days = DaysPerWeek
	}

	t := time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 23, 59, 59, 0, c.Location())
	return t.AddDate(0, 0, int(DaysPerWeek-days))
}

// StartOfDay 本日开始时间
func (c STime) StartOfDay() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 0, 0, 0, 0, c.Location())
}

// EndOfDay 本日结束时间
func (c STime) EndOfDay() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), 23, 59, 59, 0, c.Location())
}

// StartOfHour 小时开始时间
func (c STime) StartOfHour() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), 0, 0, 0, c.Location())
}

// EndOfHour 小时结束时间
func (c STime) EndOfHour() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), 59, 59, 0, c.Location())
}

// StartOfMinute 分钟开始时间
func (c STime) StartOfMinute() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), 0, 0, c.Location())
}

// EndOfMinute 分钟结束时间
func (c STime) EndOfMinute() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), 59, 0, c.Location())
}

// StartOfSecond 秒开始时间
func (c STime) StartOfSecond() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 0, c.Location())
}

// EndOfSecond 秒结束时间
func (c STime) EndOfSecond() time.Time {
	return time.Date(c.Time.Year(), c.Time.Month(), c.Time.Day(), c.Time.Hour(), c.Time.Minute(), c.Time.Second(), 999999999, c.Location())
}
