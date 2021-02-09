package reservation

import (
	"strings"
	"sync"
	"time"
)

var queue *reservationQueue

func init() {
	var q []*Reservation = make([]*Reservation, 0)
	queue = new(q)
}

// ReservationQueue is synchronoized queue
type reservationQueue struct {
	queue []*Reservation
	mutex sync.Mutex
}

func new(queue []*Reservation) *reservationQueue {
	return &reservationQueue{queue: queue, mutex: sync.Mutex{}}
}

// Enqueue an item with sync
func (q *reservationQueue) Enqueue(r *Reservation) {
	q.mutex.Lock()
	q.queue = append(q.queue, r)
	q.mutex.Unlock()
}

// Dequeue an item with sync
func (q *reservationQueue) Dequeue() *Reservation {
	q.mutex.Lock()
	var el *Reservation
	if len(q.queue) != 0 {
		el = q.queue[0]
		q.queue = q.queue[1:]
	}
	q.mutex.Unlock()
	return el
}

// DequeueExpired slice queue from expired item index caculcated by due
func (q *reservationQueue) DequeueExpired(due time.Duration) {
	q.mutex.Lock()
	var to int
	now := time.Now()
	for idx, item := range q.queue {
		reservedAt := time.Unix(item.ReserveAt, 0)
		diffMin := now.Sub(reservedAt).Minutes()
		if diffMin > due.Minutes() {
			to = idx
		}
	}
	q.queue = q.queue[to:]
	q.mutex.Unlock()
}

func (q *reservationQueue) String() string {
	s := make([]string, 0)
	for _, item := range q.queue {
		s = append(s, item.String())
	}
	return strings.Join(s, "\n")
}
