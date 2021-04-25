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
	t.Log(a, b)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)

	t.Log(math.MaxInt8)
	t.Log(math.MaxInt16)
	t.Log(math.MaxInt32)
	t.Log(math.MaxInt64)
}

func TestPoint(t *testing.T) {
	a := 2
	ptr := &a
	// ptr = ptr + 1
	t.Log(a, ptr)
	t.Logf("%T %T", a, ptr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	t.Log(s == "")
}
