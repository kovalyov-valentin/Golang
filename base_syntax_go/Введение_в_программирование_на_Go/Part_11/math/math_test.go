// Напишите серию тестов для функций Min и Max из предыдущей главы.

// package math

// import "testing"


// type testpair struct {
// 	values []int8
// 	result int8
// }
// var min_tests = []testpair {
// 	{[]int8{1,2,3,4,5}, 5},
// 	{[]int8{2,2,2,2,2}, 2},
// 	{[]int8{1,2,3,3,3}, 3},
// }
// var max_tests = []testpair {
// 	{[]int8{1,2,3,4,5}, 5},
// 	{[]int8{2,2,2,2,2}, 2},
// 	{[]int8{1,2,3,3,3}, 3},
// }
// func TestMin(t *testing.T) {
// 	for _, pair := range min_tests {
// 		value := Min(pair.values)
// 		if value != pair.result {
// 			t.Error(
// 				"For", pair.values,
// 				"expected", pair.result,
// 				"instead got", value,
// 			)
// 		}
// 	}
// }
// func TestMax(t *testing.T) {
// 	for _, pair := range max_tests {
// 		value := Max(pair.values)
// 		if value != pair.result {
// 			t.Error(
// 				"For", pair.values,
// 				"expected", pair.result,
// 				"instead got", value,
// 			)
// 		}
// 	}
// }
// func Min(xs []int8) int8 {
// 	var min = math.MaxInt8
// 	for _, i := range xs {
// 		if i < min {
// 			min = i
// 		}
// 	}
// 	return min
// }
// func Max(xs []int8) int8 {
// 	var max int8
// 	max = math.MinInt8
// 	for _, i := range xs {
// 		if i > max {
// 			max = i
// 		}
// 	}
// 	return max
// }
package math

import "testing"

type testpair struct {
	values []float64
	res    float64
}

var minTests = []testpair{
	{[]float64{1, 2, 3}, 1},
	{[]float64{2, 1, 0, -1, -2}, -2},
	{[]float64{-0.1, 0.1}, -0.1},
}

var maxTests = []testpair{
	{[]float64{1, 2, 3}, 3},
	{[]float64{-2, -1, 0, 1, 2}, 2},
	{[]float64{-0.1, 0.1}, 0.1},
}
func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.res {
			t.Error(
				"For", pair.values,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.res {
			t.Error(
				"For", pair.values,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}
func Max(xs []float64) float64 {
	max := xs[0]
	for _, val := range xs[1:] {
		if val > max {
			max = val
		}
	}
	return max
}
func Min(xs []float64) float64 {
	min := xs[0]
	for _, val := range xs[1:] {
		if val < min {
			min = val
		}
	}
	return min
}

//Пока что сам я это написать не смогу