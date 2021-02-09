package reservation

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := queue
	l := 10
	for i := 0; i < l; i++ {
		r := New(fmt.Sprintf("p:%d", i))
		offset := int64(time.Minute/time.Second) * 5 * int64(l-i)
		r.ReserveAt = r.ReserveAt - offset
		q.Enqueue(r)
	}

	t.Logf("%s", q.String())

	q.DequeueExpired(time.Minute * 30)

	t.Logf("%s", q.String())
}
