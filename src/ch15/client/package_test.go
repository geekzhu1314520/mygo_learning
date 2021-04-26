package client

import (
	"testing"

	"sankuai.com/m/src/ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(3))
	// t.Log(series.square(3))
}
