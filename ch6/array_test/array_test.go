package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[1])
	arr1 := [...]int{1, 2, 3, 4}
	t.Log(arr1[1])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}
}
