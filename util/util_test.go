package util

import (
	"strconv"
	"testing"
)

func TestSumInt(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 {
		t.Fatal("sum int failed")
	}
}

func TestSumFloat(t *testing.T) {
	if Sum([]float64{1.5, 2.5}) != 4.0 {
		t.Fatal("sum float failed")
	}
}

func TestMap(t *testing.T) {
	res := Map([]int{1, 2}, func(x int) string {
		return strconv.Itoa(x)
	})

	if res[0] != "1" || res[1] != "2" {
		t.Fatal("map failed")
	}
}

func TestFilter(t *testing.T) {
	res := Filter([]int{5, 10, 20}, func(x int) bool {
		return x >= 10
	})

	if len(res) != 2 {
		t.Fatal("filter failed")
	}
}
