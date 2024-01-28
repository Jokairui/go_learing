package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestMax(t *testing.T) {
	t.Log(math.MaxInt16)
	t.Log(math.MaxInt32)
	t.Log(math.MaxInt64)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// go 不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
