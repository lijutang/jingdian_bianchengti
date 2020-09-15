package main

import "fmt"

func equalx(x,y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
x := []string{"a", "b", "c", "d"}
y := []string{"a", "d", "p", "d"}
z := []string{"a", "b", "c", "d"}
fmt.Println(equalx(x, z))
fmt.Println(equalx(y, z))
}
//结果：
//true
//false