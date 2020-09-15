package main

import "fmt"

func reverse(s []int) []int {
	for i:=0; i < int(len(s)/2) ;i++{
		li := len(s) - i -1
		fmt.Println(i,"<=>",li)
		s[i],s[li] = s[li],s[i]
	}
	return s
}


func main() {
	var s []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reverse(s))
}

