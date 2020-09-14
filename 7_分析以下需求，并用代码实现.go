package main

import "fmt"

/*
	(1)根据工龄(整数)给员工涨工资(整数),
	工龄和基本工资通过键盘录入
	(2)涨工资的条件如下：
	[10-15) +5000
	[5-10) +2500
	[3~5) +1000
	[1~3) +500
	[0~1) +200
	(3)如果用户输入的工龄为10，基本工资为3000，
	程序运行后打印格式"您目前工作了10年，基本工资为 3000元,
	应涨工资 5000元,涨后工资 8000元"
*/

func main() {
	var n,salary,sum,a int
	var err error
	fmt.Println("请输入您的工龄")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("请输入您的基本工资")
	_, err = fmt.Scanln(&salary)
	if err != nil {
		fmt.Println(err)
	}
	if n >= 0 && n < 1 {
		a = 200
	} else if n >= 1 && n < 3 {
		a = 500
	} else if n >= 3 && n < 5 {
		a = 1000
	} else if n >= 5 && n < 10 {
		a = 2500
	} else if n >= 10 && n < 15 {
		a = 5000
	}
	sum = salary + a
	fmt.Printf("您目前工作了%d年，基本工资为%d元,应涨工资%d元,涨后工资%d元", n, salary, a, sum)
}