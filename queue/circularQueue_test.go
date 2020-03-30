package queue

import (
	"fmt"
	"testing"
)

var (
	err   error
	queue *CircularQueue
)

func TestCircularQueue_AddQueue(t *testing.T) {
	queue, err = NewCircularQueue(100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 200; i++ {
		t.Run("test", func(t *testing.T) {
			queue.In(i)
		})
	}
}

func TestCircularQueue_OutQueue(t *testing.T) {
	queue, err = NewCircularQueue(100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 200; i++ {
		t.Run("test", func(t *testing.T) {
			fmt.Println(queue.Out())
		})
	}
}
