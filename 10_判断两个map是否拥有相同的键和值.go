package main

import "fmt"

func equal(x,y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k,xValue := range x {//使用!ok来区分元素不存在和元素存在但值为0的情况
		if yValue,ok :=y[k]; !ok ||yValue != xValue {
			return false
		}
	}
	return true
}
func main() {
	x := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
	}
	y := map[string]int{
		"B": 1,
		"C": 2,
		"D": 3,
	}
	z := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
	}
	fmt.Println(equal(x, y))
	fmt.Println(equal(x, z))
}





