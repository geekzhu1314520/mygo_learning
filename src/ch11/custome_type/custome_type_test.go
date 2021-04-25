package custome_test

import (
	"fmt"
	"testing"
	"time"
)

type intConv func(op int) int

func timeSpent(inner intConv) intConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}
