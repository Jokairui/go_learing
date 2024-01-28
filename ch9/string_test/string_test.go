package string_test

import "testing"

func TestRune(t *testing.T) {
	s := "中华人民共和国"
	t.Log(len(s))
	for i, c := range s {
		t.Logf("%[1]c %[1]x %[1]d index : %[2]d ", c, i)
	}
	r := []rune(s)
	t.Log(len(r))
	for _, c := range r {
		t.Logf("%[1]c %[1]x %[1]d ", c)
	}
}

func TestUnicode(t *testing.T) {
	s := "中华人民共和国"
	t.Log(len(s))
	for _, z := range s {
		t.Logf("%[1]c %[1]x %[1]d ", z)
	}
}

func TestStringEquals(t *testing.T) {
	s1 := "中华人民共和国"
	s2 := "中华人民共和国"
	t.Log(s1 == s2)
}
