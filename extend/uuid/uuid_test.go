package sUuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Next())
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Next()
	}
}
