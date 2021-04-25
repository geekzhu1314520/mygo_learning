package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[2])

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3}

	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 5}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for idx, e := range arr2 {
		t.Log(idx, e)
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr2[1:])
	t.Log(arr2[:3])
	t.Log(arr2[1:3])
	t.Log(arr2[:])
}
