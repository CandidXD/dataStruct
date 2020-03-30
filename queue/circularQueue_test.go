package main

import (
	"fmt"
	"testing"
)

func TestCircularQueue_AddQueue(t *testing.T) {
	queue, err := NewCircularQueue(100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 200; i++ {
		t.Run("", func(t *testing.T) {
			queue.In(i)
		})
	}
}

func TestCircularQueue_OutQueue(t *testing.T) {
	queue, err := NewCircularQueue(100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 200; i++ {
		t.Run("", func(t *testing.T) {
			fmt.Println(queue.Out())
		})
	}
}
