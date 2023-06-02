package sTime

import "time"

func (c *STime) SetTimezone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	c.In(loc)
	return nil
}

// SetYear 设置年
func (c STime) SetYear(year int) STime {
	c.Time = time.Date(year, c.Time.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetMonth 设置月
func (c STime) SetMonth(month int) STime {
	c.Time = time.Date(c.Year(), time.Month(month), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetDay 设置日
func (c STime) SetDay(day int) STime {
	c.Time = time.Date(c.Year(), c.Time.Month(), day, c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetHour 设置时
func (c STime) SetHour(hour int) STime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), hour, c.Minute(), c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetMinute 设置分
func (c STime) SetMinute(minute int) STime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), minute, c.Second(), c.Nanosecond(), c.Location())
	return c
}

// SetSecond 设置秒
func (c STime) SetSecond(second int) STime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), c.Minute(), second, c.Nanosecond(), c.Location())
	return c
}

// SetNanoSecond 设置纳秒
func (c STime) SetNanoSecond(nanoSecond int) STime {
	c.Time = time.Date(c.Year(), c.Time.Month(), c.Day(), c.Hour(), c.Minute(), c.Second(), nanoSecond, c.Location())
	return c
}
