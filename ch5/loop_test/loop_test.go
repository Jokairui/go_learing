package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
	for {
		t.Log(n)
		n++
		if n > 10 {
			break
		}
	}
}
