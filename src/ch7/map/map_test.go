package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 4
	t.Log(m2[4])
	t.Log(len(m2))

	m3 := make(map[int]int, 10)
	t.Log(len(m3))
}

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is=%d", v)
	} else {
		t.Log("Key 3's value not exist.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"wang": 1, "zhu": 4, "feng": 8}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
