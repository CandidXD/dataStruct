package queue

import (
	"fmt"
	"testing"
)

var (
	queue *CircularQueue
)

func TestCircularQueue_AddQueue(t *testing.T) {
	queue = NewCircularQueue(100)
	for i := 0; i < 200; i++ {
		t.Run("test", func(t *testing.T) {
			queue.In(i)
		})
	}
}

func TestCircularQueue_OutQueue(t *testing.T) {
	queue = NewCircularQueue(100)

	for i := 0; i < 200; i++ {
		t.Run("test", func(t *testing.T) {
			fmt.Println(queue.Out())
		})
	}
}
