package main

import "fmt"

/*
	猴子第一天摘了若干个桃子，当即吃了一半，还不解馋，
	又多吃了一个；第二天，吃剩下的桃子的一半，还不过瘾，
	又多吃了一个；以后每天都吃前一天剩下的一半多一个，
	到第10天想再吃时，只剩下一个桃子了。问第一天共摘了多少个桃子？
	第10天只有一个桃子；
	第9天的桃子数 = （第10天桃子数+1）*；
	规律：第n天的桃子数：peach (n) = (peach(n+1) + 1) * 2
*/




func f(n int ) int {
	if n == 1 {
		return 1 //最后一天只有一个桃子
	} else {
		return 2*f(n-1)+2
	}
}

func eatPeach(n int64) int64 {
	if n > 10 || n < 1 {
		fmt.Println("输入数据有误,请重新输入")
	}
	if n == 10 {
		return 1
	} else {
		return (eatPeach(n+1) + 1) * 2
	}
}

func t(){
	x :=1
	for i :=10;i>=1;i--{
		fmt.Printf("第%d天的桃子数：%d\n",i,x);
		x=2*(x+1)
	}
}



func main() {
	t()
	//sum := f(10)//十天前有sum个桃子
	//fmt.Printf("第一天摘了%d个桃子", sum)
}
//结果：
//第一天摘了1534个桃子
