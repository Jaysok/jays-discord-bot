package reservation

import (
	"testing"
	"time"
)

func TestReservation(t *testing.T) {
	r := New("jaysok")

	r.ReserveAt = r.ReserveAt - int64(time.Minute*35/time.Second)
	r2 := NewWithMemo("jaysok", "안오면 그냥 ㄱ")
	t.Logf("%s", r.String())
	t.Logf("%s", r2.String())
}
