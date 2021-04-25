package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 2, 3, 5}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
	t.Log(a == e)
}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	a = a &^ Executeable
	t.Log(Readable&a == Readable, Writeable&a == Writeable, Executeable&a == Executeable)
}
