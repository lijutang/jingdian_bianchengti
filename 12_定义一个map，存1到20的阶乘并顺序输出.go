package main

import (
	"fmt"
	"sort"
)

//法一、

func sort1() {
	m := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i

		}
		fmt.Println(i, "的阶乘是", m[i])
	}
}

//法二、
func sort2(){
	m := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i

		}
	}
	arr := make([]int,0)
	for k,_ := range m {
		arr = append(arr,k)
	}
	sort.Ints(arr)
	for i :=0;i<=len(arr)-1;i++ {
		fmt.Println(arr[i],"的阶乘是",m[arr[i]])
	}

}

func main() {
	//sort1()
	sort2()
}


