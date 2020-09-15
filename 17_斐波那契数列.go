package main

import "fmt"

/*
	斐波那契数列指的是这样一个数列 1, 1, 2, 3, 5, 8, 13, 21, 34, 55…
	这个数列从第3项开始，每一项都等于前两项之和
*/


func isFibonacciSequence(i int) int {
	if i <= 1 {
		return 1
	}
	return isFibonacciSequence(i-1) + isFibonacciSequence(i-2)
}

func main() {
	var i int = 10
	var j int =3

	fmt.Println("第", i+1, "项的值为", isFibonacciSequence(i))
	fmt.Println("第", j+1, "项的值为", isFibonacciSequence(j))
}