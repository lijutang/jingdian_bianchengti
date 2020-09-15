package main

import (
	"fmt"
	"sort"
)

//写一个函数实现顺序遍历map中的元素

func sortMap(student map[string]int) {
	var names []string
	for name := range student {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s的年龄是%d\n", name, student[name])
	}
}
func main() {
	student := map[string]int{
		"lisa":     17,
		"bob":      20,
		"victoria": 24,
		"sabit":    40,
	}
	sortMap(student)
}
