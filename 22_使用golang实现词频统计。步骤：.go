package main

import (
	"btctrans/common/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
	（1）从文件中读取一篇文章。
	（2）统计词频，按单词出现的频率从大到小进行排序
	（3）写入到文件中。
*/
func getWordFrequency(readFilePath string, writeFilePath string) {
	var fileText string
	var wordFrequencyMap = make(map[string]int)
	//读取文件
	fileData, err := ioutil.ReadFile(readFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fileText = string(fileData)
	//去掉分隔符
	f := func(c rune) bool {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			return true
		}
		return false
	}
	arr := strings.FieldsFunc(fileText, f)
	//如果字典里有该单词则加1，否则添加入字典赋值为1
	for _, v := range arr {
		if _, ok := wordFrequencyMap[v]; ok {
			wordFrequencyMap[v] = wordFrequencyMap[v] + 1
		} else {
			wordFrequencyMap[v] = 1
		}
	}
	//fmt.Println(wordFrequencyMap)
	//按照单词出现的频率排序
	type wordFrequencyNum struct {
		Word string
		Num  int
	}
	var lstWordFrequencyNum []wordFrequencyNum
	for k, v := range wordFrequencyMap {
		lstWordFrequencyNum = append(lstWordFrequencyNum, wordFrequencyNum{k, v})
	}
	sort.Slice(lstWordFrequencyNum, func(i, j int) bool {
		return lstWordFrequencyNum[i].Num > lstWordFrequencyNum[j].Num
	})
	fmt.Println("按照单词出现频率由高到低排序", lstWordFrequencyNum)
	//写入文件
	var jsonBytes []byte
	var arrJsonBytes string
	for _, v := range lstWordFrequencyNum {
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		arrJsonBytes = arrJsonBytes + string(jsonBytes)
	}
	err = ioutil.WriteFile(writeFilePath, []byte(arrJsonBytes), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	readFilePath := "C:\\Users\\zyq\\go\\src\\yj\\artical.txt"//读取的文件路径
	writeFilePath := "C:\\Users\\zyq\\go\\src\\yj\\wordFrequency.txt"//写入的文件路径
	getWordFrequency(readFilePath, writeFilePath)
}