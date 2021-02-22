package stack

import "log"

type Queue struct {
	Front, Rear, MaxQueue int
	Elements              [MaxSize]T
}

//CreateQueue 构造一个空队列
func CreateQueue(q *Queue, maxSize int) {
	q.Front = 0
	q.Rear = 0
	q.MaxQueue = maxSize
}

//IsEmpty 队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.Front == q.Rear
}

//IsFull 队列是否已满
func (q *Queue) IsFull() bool {
	return (q.Rear+1)%q.MaxQueue == q.Front
}

//Append 入队
func (q *Queue) Append(x T) {
	if q.IsFull() {
		log.Fatalf("%s", "Overflow")
	}
	q.Rear = (q.Rear + 1) % q.MaxQueue
	q.Elements[q.Rear] = x
}

//Serve 从队头删除元素
func (q *Queue) Serve() {
	if q.IsEmpty() {
		log.Fatalf("%s", "Underflow")
	}
	q.Front = (q.Front + 1) % q.MaxQueue
}

//QueueFront 从对头返回元素
func (q *Queue) QueueFront(x *T) {
	if q.IsEmpty() {
		log.Fatalf("%s", "Underflow")
	}
	*x = q.Elements[(q.Front+1)%q.MaxQueue]
}
