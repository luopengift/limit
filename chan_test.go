package limit

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel(t *testing.T) {
	ch := NewChannel(10)
	//启动协程查看信息
	for i := range []int{0, 1, 2, 3, 4, 5} {
		err := ch.Run(func() error {
			return nil
		})
		fmt.Println(i, err)
	}
	for i := 0; i < 10; i++ {
		ch.Put(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("err:", ch.Close())
}

func Benchmark_Channel(b *testing.B) {
	ch := NewChannel(1000)
	/*go func() {
		for {
			if _,ok := ch.Get(); !ok {
				return
			}
		}
	}()*/
	for i := 0; i < b.N; i++ {
		ch.Put("11111111111111111111111111111asdfasdfadfl;kj11111111111111111111111")
		ch.Get()
	}
	ch.Close()
}
