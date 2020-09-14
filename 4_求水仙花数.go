package main

import "fmt"

/*
	水仙花数是指一个 3 位数，它的每个位上的数字的
	3次幂之和等于它本身（例如：1^3 + 5^3+ 3^3 = 153）
*/

func isNarcissisticNum(num int) bool {
	a := num / 100       //分离出百位a
	b := (num / 10) % 10 //分离处十位b
	c := num % 10        //分离出个位c
	result := a*a*a + b*b*b + c*c*c
	if num == result {
		return true
	}
	return false
}

func main() {
	for i := 100; i < 1000; i++ {
		if isNarcissisticNum(i) {
			fmt.Println("水仙花数有：", i)
		}
	}
}
//结果：
//水仙花数有： 153
//水仙花数有： 370
//水仙花数有： 371
//水仙花数有： 407