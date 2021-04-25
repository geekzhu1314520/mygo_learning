package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	// s[1] = '3' cannot assign to s[1] (strings are immutable)

	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "\xE4\xB8\xA5\xFF"
	t.Log(s)

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中的unicode %x", c[0])
	t.Logf("中的UTF-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]d", c)
	}
}

func TestStringToRune2(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c, %[1]x", c)
	}
}
