package try_constant

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executeable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	// a := 7
	a := 1
	t.Log(Readable&a == Readable, Writeable&a == Writeable, Executeable&a == Executeable)
}
