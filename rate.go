package limit

import (
	"fmt"
	"sync"
	"time"
)

// Limit limit
type Limit struct {
	ID       string
	mux      *sync.Mutex
	rate     int           //速率
	interval int           //时间间隔
	count    int           //当前计数
	endTime  time.Time     //每个时间区间的结束时间
	duration time.Duration //时间间隔

}

// NewLimit NewLimit
func NewLimit(name string, rate, interval int) *Limit {
	limit := &Limit{
		ID:  name,
		mux: new(sync.Mutex),
	}
	limit.Set(rate, interval)
	limit.Reset()
	limit.reset()
	return limit
}

func (l *Limit) String() string {
	return fmt.Sprintf("endtime:%v", l.endTime)
}

// Set Limit rate and interval.
func (l *Limit) Set(rate, interval int) {
	l.rate = rate
	l.interval = interval
	l.duration = time.Duration(l.interval) * time.Second
}

// Reset Limit count
func (l *Limit) Reset() {
	l.endTime = time.Now().Add(l.duration)
	l.count = 0
}

func (l *Limit) reset() {
	go func() {
		t := time.NewTimer(l.duration)
		for {
			select {
			case <-t.C:
				l.Reset()
				t.Reset(l.duration)
			}
		}
	}()
}

// IsLimit is limit
func (l *Limit) IsLimit() (bool, time.Duration) {
	if l.count < l.rate {
		l.count++
		return false, 0
	}
	return true, l.endTime.Sub(time.Now())
}

// Wait wait
func (l *Limit) Wait(fun func() error) error {
	for {
		if ok, d := l.IsLimit(); ok {
			time.Sleep(d)
			continue
		}
		return fun()
	}
}

// Do do
func (l *Limit) Do(fun func() error) error {
	if ok, _ := l.IsLimit(); ok {
		return fmt.Errorf("limit")
	}
	return fun()
}
