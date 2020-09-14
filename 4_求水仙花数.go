package main

import "fmt"

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