package onethread

import (
	"go-server/core/base"
)

var MaxTaskNum = 1000 // 每帧处理任务数量
var Instance = &OneThreadSyncContext{}

type OneThreadSyncContext struct {
	TakeQueue *base.Queue
}

func NewOneThreadSyncContext() *OneThreadSyncContext {
	return &OneThreadSyncContext{
		TakeQueue: &base.Queue{},
	}
}

func (context *OneThreadSyncContext) Post(fn func()) {
	context.TakeQueue.Enqueue(fn)
}

func (context *OneThreadSyncContext) Update() {
	for i := 0; i < MaxTaskNum; i++ {
		Instance.TakeQueue.Dequeue()
	}
}
