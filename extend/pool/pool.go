package sPool

import (
	"bytes"
	"sync"
)

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func GetBuffer(data []byte) *bytes.Buffer {
	buffer := bufferPool.Get().(*bytes.Buffer)
	if data != nil {
		buffer.Write(data)
	}
	return buffer
}

func PutBuffer(buffer *bytes.Buffer) {
	if buffer != nil {
		buffer.Reset()
		bufferPool.Put(buffer)
	}
}
