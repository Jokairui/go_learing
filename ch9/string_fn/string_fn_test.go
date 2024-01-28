package string_fn

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string = "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
