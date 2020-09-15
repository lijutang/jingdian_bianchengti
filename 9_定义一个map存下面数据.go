package main

import (
	"database/sql"
	"fmt"
)

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
	//使用键输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
	//查看元素在集合中是否存在
	capital,ok := countryCapitalMap["American"]//如果确定是真实的,则存在,否则不存在
	if ok {
		fmt.Println("American的首都是",capital)
	} else {
		fmt.Println("merican的首都不存在")
	}
	//用户名：密码＠tcp(地址：3306)/是数据库名
	db,err := sql.Open("mysql","root:123456@tcp(192.168.1.15)/test")
	if err != nil {
		fmt.Println(err)
	}
	//网数据库中插如数据
	for k,v := range countryCapitalMap {
		result,err := db.Exec("INSERT  INTO countryCapital(country,capital)values(?,?) ",k,v)
		if err != nil {
			fmt.Println(result,err)
		}
	}
}
