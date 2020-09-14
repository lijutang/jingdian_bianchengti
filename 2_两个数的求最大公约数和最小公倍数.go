package main

import "fmt"

func getMaximumCommonDivisor(a,b int) int {
	//获取最大公约数
	for a!=b {
		if a>b {
			a= a-b
		} else if a <b {
			b = b-a

		}
	}
	return a
}


func main() {
	var a, b = 24, 10
	num := getMaximumCommonDivisor(a, b)
	fmt.Println("a,b的最大公约数是：", num)
	//求最小公倍数相对来说就比较简单了。只需要先求出最大公约数。用两个数的乘积除以最大公约数即可
	fmt.Println("a,b的最小公倍数是：", a*b/num)
}
//结果：
//a,b的最大公约数是： 2
//a,b的最小公倍数是： 120