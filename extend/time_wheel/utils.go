package sTimeWheel

import (
	sTime "go-server/extend/time"
	"sync"
	"sync/atomic"
)

// truncate returns the result of rounding x toward zero to a multiple of m.
// If m <= 0, Truncate returns x unchanged.
func truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}

// TimeToMS returns an integer number, which represents t in milliseconds.
func TimeToMS(sTime sTime.STime) int64 {
	return sTime.ToMillisecond()
}

// MSToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func MSToTime(t int64) sTime.STime {
	return sTime.NewMillisecond(t)
}

type waitGroupWrapper struct {
	sync.WaitGroup
}

func (w *waitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}

var _nextId uint64

func NextId() uint64 {
	return atomic.AddUint64(&_nextId, 1)
}
