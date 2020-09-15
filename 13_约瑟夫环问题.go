package main

import "fmt"

//编号为 1-N 的 N 个士兵围坐在一起形成一个圆圈，
//从编号为 1 的士兵开始依次报数（1，2，3…这样依次报）
//，数到 k 的 士兵会被杀掉出列，之后的士兵再从 1 开始报数。
//直到最后剩下一士兵，求这个士兵的编号。

/*
实现一
	递推方法,推算出公式:f(n)=[f(n-1,m)+m]%m (n>1) 剑指offer 62题
	时间复杂度O(n)
*/
func Jos1(number, index int) int {
	//处理边界
	if number < 0 || index < 0 {
		return -1
	}
	last := 0
	for i := 2; i <= number; i ++ {
		last = (last + index) % i
	}
	return last+1
}
/*
实现二
	模拟一个环, 直接剔出第一个人,然后再验证是否符合条件,不符合追加到环尾部.直到环中最后一个元素
	时间复杂度O(n*m)
	空间复杂度O(n)
*/

//模拟一个环, 直接剔出第一个人,然后再验证是否符合条件,
// 不符合追加到环尾部.直到环中最后一个元素
func Jos2(number, index int) int {
	//处理边界
	if number < 0 || index < 0 {
		return -1
	}
	//申请一个切片,模拟一个环.一个位置一个人.
	people := make([]int, number)
	for i := 0; i < number; i ++ {
		people[i] = i+1
	}
	i := 0 //编号从0开始
	for len(people) > 1 { //只要切片里的人数还大于1个人则继续数.
		i ++ //正式开始数
		head := people[0] //直接出列第一个人
		people = people[1:] //将从环中删除
		if i % index != 0 { //查看此人是不是中招的人,不是则重新加入环后面.
			people = append(people, head)
		}
	}
	return people[0]
}

/*
实现三
	模拟一个环, 如果是目标则对环中的元素标记为0, 直到环中只有一个元素大于0
	时间复杂度O(n*m)
	空间复杂度O(n)
 */
//模拟一个环, 如果是目标则对环中的元素标记为0, 直到环中只有一个元素大于0
func Jos3(number, index int) int {
	if number < 0 || index < 0 {
		return -1
	}
	//申请一个切片,模拟一个环.一个位置一个人.
	people := make([]int, number)
	for i := 0; i < number; i ++ {
		people[i] = i+1
	}
	j := 0 //数组指针, 向后走
	l := 0 //记录退出的人数
	k := 0 //报数,1,2,3
	for l < number -1 {
		if people[j] != 0 { //未出局的人就开始数.
			k++
		}
		if k == index {
			people[j] = 0 //出局的人
			k = 0 //重新算
			l++ //记录已经出局1人
		}
		j ++ //继续移动数组指针
		if j == number {//如果移动到数组尾则从头算开始移动
			j = 0
		}
	}
	//大于0的则是目标值
	for i := 0; i < number; i ++ {
		if people[i] > 0 {
			return people[i]
		}
	}
	return -1
}

//递归
func josephus(n int,k int) int {
	if n == 1 {
		return n
	}
	return (josephus(n-1,k)+ k-1)%n +1
}

func main() {
	n := 5
	k := 2
	//res := Jos1(n, k)
	//res := Jos2(n, k)
	res := Jos3(n, k)
	//res := josephus(n, k)
	fmt.Println(res)
}






