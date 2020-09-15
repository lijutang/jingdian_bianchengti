package main

import "fmt"

/*
翻转下面这个二维数组
	a = [3][4]int{
	{0, 1, 2, 3} ,
	{4, 5, 6, 7} ,
	{8, 9, 10, 11},
	}
*/

func transpose(a [][]int) [][]int {
	//只有一行则不是二维数组
	if len(a[0]) == 0 {
		return nil
	}
	row := len(a[0]) //arr的row
	col := len(a)//arr的col
	arr := make([][]int ,row,row)
	for i:=0; i<=row-1;i++ {
		arr[i] = make([]int,col)
	}
	for k := col - 1; k >= 0; k-- {
		for y := row - 1; y >= 0; y-- {
			arr[y][k] = a[k][y]

		}

	}

	fmt.Println(arr)
	return arr
}




func main() {
	a := [][]int{
		{0, 1, 2, 3} ,
		{4, 5, 6, 7} ,
		{8, 9, 10, 11},
	}
	transpose(a)

}