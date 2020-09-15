package main

import "fmt"

func countUnicode(s string) {
	var letter int //字母个数
	var num int    //数字个数
	var space int  //空格个数
	var other int  //其他个数
	for i := 0;i<len(s);i++ {
		if (s[i]>='a' && s[i]<= 'z') || (s[i]>='A' && s[i]<='Z') {
			letter++
		} else if s[i] >='0' && s[i]<='9' {
			num++
		} else if s[i]==' ' {
			space++
		} else {
			other++
		}
	}
	fmt.Printf("字母的个数为%d\n", letter)
	fmt.Printf("数字的个数为%d\n", num)
	fmt.Printf("空格的个数为%d\n", space)
	fmt.Printf("其他字符的个数为%d\n", other)

}

func main() {
	s:= "A, B, 44...,5 Z、a, b, ..., &&z  "
	countUnicode(s)
}