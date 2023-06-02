package sTime

import (
	cstring "go-server/extend/string"
	"time"
)

// ToSecond 输出秒级时间戳
func (c STime) ToSecond() int64 {
	return c.Unix()
}

// ToMillisecond 输出毫秒级时间戳
func (c STime) ToMillisecond() int64 {
	return c.Time.UnixNano() / int64(time.Millisecond)
}

func (c STime) ToMillisecondString() string {
	t := c.ToMillisecond()
	return cstring.ToString(t)
}

// ToMicrosecond 输出微秒级时间戳
func (c STime) ToMicrosecond() int64 {
	return c.UnixNano() / int64(time.Microsecond)
}

// ToNanosecond 输出纳秒级时间戳
func (c STime) ToNanosecond() int64 {
	return c.UnixNano()
}

// ToDateMillisecondFormat 2023-04-10 12:26:57.420
func (c STime) ToDateMillisecondFormat() string {
	return c.Format(DateTimeMillisecondFormat)
}

// ToDateTimeFormat 2006-01-02 15:04:05
func (c STime) ToDateTimeFormat() string {
	return c.Format(DateTimeFormat)
}

// ToDateFormat 2006-01-02
func (c STime) ToDateFormat() string {
	return c.Format(DateFormat)
}

// ToTimeFormat 15:04:05
func (c STime) ToTimeFormat() string {
	return c.Format(TimeFormat)
}

// ToShortDateTimeFormat 20060102150405
func (c STime) ToShortDateTimeFormat() string {
	return c.Format(ShortDateTimeFormat)
}

// ToShortDateFormat 20060102
func (c STime) ToShortDateFormat() string {
	return c.Format(ShortDateFormat)
}

// ToShortIntDateFormat 20060102
func (c STime) ToShortIntDateFormat() int32 {
	strDate := c.ToShortDateFormat()
	intDate, _ := cstring.ToInt32(strDate)
	return intDate
}

// ToShortTimeFormat 150405
func (c STime) ToShortTimeFormat() string {
	return c.Format(ShortTimeFormat)
}
