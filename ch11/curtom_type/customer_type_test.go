package curtom_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpend(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	tsSf := timeSpend(slowFunc)
	t.Log(tsSf(10))
}
