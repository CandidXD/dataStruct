package main

import (
	"fmt"
)

// 队列（根据no从小到大排序）
type QueueNode struct {
	no      int
	element int
	next    *QueueNode
}

// In 入队
func (q *QueueNode) In(no, element int) {
	if q.next != nil && q.next.no > no {
		temp := q.next
		q.next = new(QueueNode)
		q.next.no = no
		q.next.element = element
		q.next.next = temp
	} else if q.next == nil {
		q.next = new(QueueNode)
		q.next.no = no
		q.next.element = element
	} else {
		q.next.In(no, element)
	}
}

// Out 出队
func (q *QueueNode) Out() int {
	if q.next == nil {
		fmt.Println("队列为空")
		return 0
	}

	out := q.next.element
	q.next = q.next.next
	return out
}

// NewQueue 初始化队列
func NewQueue() *QueueNode {
	q := new(QueueNode)
	return q
}

// ShowQueue 打印队列
func (q *QueueNode) ShowQueue() {
	if q.next == nil {
		fmt.Println()
		return
	}
	fmt.Printf("No:%d Element:%d => ", q.next.no, q.next.element)
	q.next.ShowQueue()
}
