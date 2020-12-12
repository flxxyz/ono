package queue

import "github.com/oleiade/lane"

var queue *lane.Queue

func init() {
	queue = lane.NewQueue()
}

// Enqueue 进队列
func Enqueue(item interface{}) {
	queue.Enqueue(item)
}

// Dequeue 出队列
func Dequeue() interface{} {
	return queue.Dequeue()
}

// Head 返回队列第一项
func Head() interface{} {
	return queue.Head()
}
