package main

import (
	"fmt"
	"testing"
)

func TestQueueNode_In(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 100; i++ {
		t.Run("", func(t *testing.T) {
			queue.In(i, i)
		})
	}
}

func TestQueueNode_Out(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 100; i++ {
		t.Run("", func(t *testing.T) {
			fmt.Println(queue.Out())
		})
	}
}
