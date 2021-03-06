package limit

import (
	"fmt"
	"time"

	"github.com/luopengift/log"
)

// Chan channel
type Chan chan interface{} //使用chan堵塞控制队列的最大数量

// NewChannel channel
func NewChannel(max int) Chan {
	return make(chan interface{}, max)
}

// Close channel
func (c Chan) Close() error {
	for i := 0; i < 10; i++ {
		if len(c) == 0 {
			close(c)
			return nil
		}
		time.Sleep(10 * time.Millisecond)
	}
	return fmt.Errorf("closed ch failed! ch is not empty, len is %d", c.Len())
}

//Put 往管道中写数据
func (c Chan) Put(v interface{}) error {
	c <- v
	return nil
}

// Get 从管道中读数据
func (c Chan) Get() (interface{}, bool) {
	v, ok := <-c
	return v, ok
}

// Add 往管道中放入一个标记，记录活跃数值
func (c Chan) Add() {
	c.Put(struct{}{})
}

// Done 从管道中取出一个标记，减少活跃数值
func (c Chan) Done() {
	c.Get()
}

// Cap cap
func (c Chan) Cap() int { return cap(c) }

// Len len
func (c Chan) Len() int { return len(c) }

// Run run
func (c *Chan) Run(fun func() error) error {
	c.Add()
	go func() {
		if err := fun(); err != nil {
			log.Errorf("%s", err)
		}
		c.Done()
	}()
	return nil
}
