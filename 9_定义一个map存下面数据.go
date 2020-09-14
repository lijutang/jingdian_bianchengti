package main

/*
	1）定义一个map存下面数据
	France 首都是 巴黎
	Italy 首都是 罗马
	Japan 首都是 东京
	India 首都是 新德里
   （2）检测American 的首都是否存在
   （3）将map的数据存入到mysql数据库

*/
func main() {
	//创建集合
	var countryCapitalMap map[string] string
	countryCapitalMap = make(map[string]string)
	// map插入key - value对,各个国家对应的首都
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
}
