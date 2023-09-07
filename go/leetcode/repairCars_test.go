package leetcode

import "testing"

type CartTestData struct {
	res   int64
	cars  int
	ranks []int
}

func TestRepairCars(t *testing.T) {
	var testData = []*CartTestData{
		&CartTestData{cars: 10, ranks: []int{4, 2, 3, 1}, res: 16},
		&CartTestData{cars: 6, ranks: []int{5, 1, 8}, res: 16},
	}
	for _, d := range testData {
		i := repairCars(d.ranks, d.cars)
		if i != d.res {
			t.Error("错误", i)
		}
	}
}
