package main

import "fmt"

/*
	回文数的概念：即是给定一个数，这个数顺读和逆读都是一样的。
	例如：121，1221是回文数，123，1231不是回文数。
*/

func isHuiwen(a string) bool {
	var i, j int
	var b bool
	for i = 0; i <= len(a)/2-1; i++ {
		j = len(a) - 1
		if a[i] != a[j] {
			b = false
		}
		b = true
		j--
	}
	return b
}

func main() {
	a := "1002332001"
	isHuiwen(a)
	fmt.Println(isHuiwen(a))
}