package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
}
