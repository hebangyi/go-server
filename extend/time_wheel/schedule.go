package sTimeWheel

import (
	sTime "go-server/extend/time"
	"time"
)

// Scheduler determines the execution plan of a task.
type Scheduler interface {
	// Next returns the next execution time after the given (previous) time.
	// It will return a zero time if no next time is scheduled.
	//
	// All times must be UTC.
	Next(sTime sTime.STime) sTime.STime
}

type EverySchedule struct {
	Interval time.Duration
}

func (s *EverySchedule) Next(prev sTime.STime) sTime.STime {
	prev.AddDuration(s.Interval)
	return prev

}

type FixedDateSchedule struct {
	Hour, Minute, Second int
}

/*func (s *FixedDateSchedule) Next(prev sTime.STime) sTime.STime {
	hour := prev.Hour()
	if s.Hour >= 0 {
		hour = s.Hour
	}

	fixedTime := time.Date(
		prev.Year(),
		prev.Month(),
		prev.Day(),
		hour,
		s.Minute,
		s.Second,
		0,
		prev.Location(),
	)

	remain := fixedTime.UnixNano() - prev.UnixNano()
	if remain > 0 {
		return prev.Add(time.Duration(remain))
	}

	if s.Hour >= 0 {
		return fixedTime.Add(24 * time.Hour)
	}

	return fixedTime.Add(time.Hour)
}
*/
