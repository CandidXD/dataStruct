package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var array []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		array = append(array, rand.Int())
	}
	for i := 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			QuickSort(&array)
		})
	}
}
