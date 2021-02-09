package reservation

import (
	"fmt"
	"strings"
	"time"
)

type Reservation struct {
	Name      string
	ReserveAt int64
	Memo      string
}

func New(name string) *Reservation {
	return &Reservation{Name: name, ReserveAt: time.Now().Unix()}
}

func NewWithMemo(name string, memo string) *Reservation {
	r := New(name)
	r.Memo = memo
	return r
}

func (r *Reservation) String() string {
	reserveTime := time.Unix(r.ReserveAt, 0)
	timeDiff := time.Now().Sub(reserveTime)
	min := timeDiff.Minutes()

	s := fmt.Sprintf("%s님이 %.1f분 전에 예약", r.Name, min)
	if r.Memo == "" {
		return s
	}
	return strings.Join([]string{s, r.Memo}, ", ")
}
