package sTimeWheel

import (
	"container/list"
	"sync/atomic"
	"unsafe"
)

// TimerTask represents a single event. When the TimerTask expires, the given
// task will be executed.
type TimerTask struct {
	id         uint64
	expiration int64 // in milliseconds
	task       func()
	// The bucket that holds the list to which this timer's element belongs.
	//
	// NOTE: This field may be updated and read concurrently,
	// through TimerTask.Stop() and Bucket.Flush().
	b       unsafe.Pointer // type: *bucket
	element *list.Element  // The timer's element.
	isAsync bool           // async execute task
}

func (t *TimerTask) ID() uint64 {
	return t.id
}

func (t *TimerTask) getBucket() *bucket {
	return (*bucket)(atomic.LoadPointer(&t.b))
}

func (t *TimerTask) setBucket(b *bucket) {
	atomic.StorePointer(&t.b, unsafe.Pointer(b))
}

// Stop prevents the TimerTask from firing. It returns true if the call
// stops the timer, false if the timer has already expired or been stopped.
//
// If the timer t has already expired and the t.task has been started in its own
// goroutine; Stop does not wait for t.task to complete before returning. If the caller
// needs to know whether t.task is completed, it must coordinate with t.task explicitly.
// TODO 这里可以优化 可以不用操作 bucket
func (t *TimerTask) Stop() bool {
	stopped := false
	for b := t.getBucket(); b != nil; b = t.getBucket() {
		// If b.Remove is called just after the timing wheel's goroutine has:
		//     1. removed t from b (through b.Flush -> b.remove)
		//     2. moved t from b to another bucket ab (through b.Flush -> b.remove and ab.Add)
		// this may fail to remove t due to the change of t's bucket.
		stopped = b.Remove(t)

		// Thus, here we re-get t's possibly new bucket (nil for case 1, or ab (non-nil) for case 2),
		// and retry until the bucket becomes nil, which indicates that t has finally been removed.
	}
	return stopped
}
