package main

import "fmt"

/*

	题目的意思是，给我们一个字符串，
	让我们求最长的无重复字符的子串，
	注意这里是子串，不是子序列，所以必须是连续的。
*/

func lengthOfLongestSubstring(s string) int {
	var left, res int
	usedChar := make(map[rune]int)
	for i, v := range s {
		if k, ok := usedChar[v]; ok && left <= k {
			left = k + 1
		} else {
			res = max(res, i-left+1)
		}

		usedChar[v] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var t string = "3456789x1234123"
	var m string = "sdkjfisf56789x1234123"
	fmt.Println(lengthOfLongestSubstring(t))
	fmt.Println(lengthOfLongestSubstring(m))
}
