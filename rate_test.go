package limit

import (
	"testing"

	"github.com/luopengift/log"
)

func Test_Limit(t *testing.T) {
	limit := NewLimit("Test", 1, 1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			limit.Do(func() error {
				log.Infof("%v", i)
				return nil
			})
		}(i)
	}
	select {}
}
