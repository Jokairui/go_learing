package once

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct{}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
