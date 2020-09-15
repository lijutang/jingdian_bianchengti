package main

import "fmt"

/*
	[问题描述】
	给定一个数组，求两个数之和=给定值sum的所有组合个数。
	【方法一】穷举法
	从数组中任意找两个数，看其和是否=sum。时间复杂度O(N^2)
--------------------------------------------------------
	【方法二】hash表法
	只需要遍历一遍数组，非常高效
	思路：定义一个map，一开始为空，不存数据，开始遍历数组，
	判断第一个元素是否有另一半已经在map中，如果有count++，没有的话，
	把第一个元素存入map；继续遍历第二个元素，
	判断第二个元素是否有另一半已经在map中，如果有count++，
	没有的话，把第二个元素存入map…遍历结束，结果也出来了
	代码：
*/

func getSumNum(arr []int, sum int) int {
	tp := make(map[int]int)//map存数组里的元素，注意，这里map初始为空
	var count int//计个数

	for i := 0; i <= len(arr)-1; i++ {//遍历数组
		//判断这个元素的“另一半”是否已经在map里
		if _, ok := tp[sum-arr[i]]; ok {//如果在，就说明数组中有这个元素的“另一半”
			count++
		} else {//如果不在，就把这个元素加入map中
			tp[arr[i]] = i
		}

	}
	return count
}

func main() {
	var brr []int = []int{1, 12, 14, 40, 56, 60, 88, 78, 99}
	var sum int = 100
	fmt.Println(getSumNum(brr, sum))
}
