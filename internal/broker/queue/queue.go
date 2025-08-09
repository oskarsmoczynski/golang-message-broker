package queue

import (
	"sync"

	"github.com/oskarsmoczynski/golang-message-broker/internal/broker/models"
)

type node struct {
	msg  models.Message
	next *node
}

type Queue struct {
	head, tail *node
	size       int
	mu         sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Enqueue(message models.Message) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newNode := &node{msg: message}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *Queue) Dequeue() (models.Message, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.head == nil {
		return models.Message{}, false
	}

	msg := q.head.msg
	q.head = q.head.next
	q.size--

	if q.head == nil {
		q.tail = nil

		// This might be redundant since size is decremented before, but I'm a fan of sanity checks
		if q.size != 0 {
			q.size = 0
		}
	}

	return msg, true
}

func (q *Queue) Empty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.size == 0
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.size
}
