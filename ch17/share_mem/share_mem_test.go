package share_mem

import (
	"sync"
	"testing"
)

func TestShareMem(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("count : ", count)

}
