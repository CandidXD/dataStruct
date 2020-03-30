package main

import (
	"errors"
	"fmt"
)

// 环形队列
type CircularQueue struct {
	array   []int // 数组
	maxSize int   // 最大长度
	front   int   // 头部下标（不包含）
	rear    int   // 尾部下标（包含）
	len     int   // 实际元素长度
}

// GetMaxSize 获取最大长度
func (q *CircularQueue) GetMaxSize() int {
	return q.maxSize
}

// GetFront 获取头部下标
func (q *CircularQueue) GetFront() int {
	return q.front
}

// GetRear 获取尾部下标
func (q *CircularQueue) GetRear() int {
	return q.rear
}

// AddQueue 入队
func (q *CircularQueue) In(element int) {
	// 判断队列是否已满
	if q.len == q.maxSize {
		fmt.Println("队列已满")
		return
	}

	// 尾部下标右移
	q.rear++

	// 尾部下标大于最大长度时，尾部标下循环
	if q.rear > q.maxSize-1 {
		q.rear = 0
	}

	// 尾部添加元素
	q.array[q.rear] = element

	// 实际长度增加
	q.len++
}

// OutQueue 出队
func (q *CircularQueue) Out() int {
	// 判断队列是否为空
	if q.len == 0 {
		fmt.Println("队列为空")
		return 0
	}

	// 获取队首元素
	out := q.array[q.front+1]

	// 实际长度减少
	q.len--
	// 获取队首元素置空
	q.array[q.front+1] = 0

	// 首部下标右移
	q.front++

	// 头部下标等于最大长度，头部下标循环
	if q.front == q.maxSize-1 {
		q.front = -1
	}

	return out
}

// ShowQueue 打印队列
func (q *CircularQueue) ShowQueue() {
	fmt.Println(q.array)
}

// NewQueue 创建队列
func NewCircularQueue(maxSize int) (circularQueue *CircularQueue, err error) {
	if maxSize < 1 {
		return nil, errors.New("队列长度不能小于1")
	}
	q := new(CircularQueue)
	q.maxSize = maxSize
	q.front = -1
	q.rear = -1
	for i := 0; i < maxSize; i++ {
		q.array = append(q.array, 0)
	}
	return q, nil
}
