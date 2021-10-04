package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// map（映射）
func main() {
	// 声明map类型：未初始化，初始值为nil
	var a map[string]int
	fmt.Println(a == nil)
	// map 初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	// map添加数据（键值对）
	a["OHIUA"] = 24
	a["SHINRIN"] = 23
	fmt.Printf("a: %#v \ntype: %T\n", a, a)

	// 声明map的同时初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b: %#v \ntype: %T\n", b, b)

	// map未初始化时为nil，无法直接操作
	// var c map[int]int
	// c[100] = 200 // c 未初始化（未申请内存），无法直接操作
	// fmt.Println(c)

	// map查询：判断键是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["YaSuo"] = 100
	scoreMap["Teemo"] = 99
	v, ok := scoreMap["Teemo"]
	if ok {
		fmt.Println("Teemo in the scoreMap:", v)
	} else {
		fmt.Println("Teemo is not in the scoreMap.")
	}

	// map遍历：for range，key或v可忽略
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}

	// map删除元素
	delete(scoreMap, "YaSuo")
	fmt.Println(scoreMap)

	// 按照指定顺序遍历map
	var heroMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("hero%02d", i)
		value := rand.Intn(100)
		heroMap[key] = value
	}
	// 1.取出key放入切片
	keys := make([]string, 0, len(heroMap))
	for k := range heroMap {
		keys = append(keys, k)
	}
	// 2.对切片中的key排序
	sort.Strings(keys)
	// 3.按排序后的key的顺序遍历heroMap
	for _, key := range keys {
		fmt.Println(key, heroMap[key])
	}

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8) // 完成切片的初始化
	mapSlice[0] = make(map[string]int, 8)    // 完成map的初始化
	mapSlice[0]["ZZYO"] = 23
	mapSlice[0]["AZE"] = 22

	for v := range mapSlice {
		for k, u := range mapSlice[v] {
			fmt.Print(k, ":", u, " ")
		}
	}
	fmt.Println()

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 完成map的初始化
	val, isOk := sliceMap["Jee"]
	if isOk {
		fmt.Println(val)
	} else {
		sliceMap["Jee"] = make([]int, 8) // 完成切片初始化
		sliceMap["Jee"][0] = 100
		sliceMap["Jee"][1] = 200
		sliceMap["Jee"][2] = 300
	}

	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	// 统计字符串中每个单词出现的次数
	var str = "how do u do"
	words := strings.Split(str, " ")
	var count = make(map[string]int, len(words))
	for _, word := range words {
		v, ok := count[word]
		if ok {
			count[word] = v + 1
		} else {
			count[word] = 1
		}
	}
	for k, v := range count {
		fmt.Println(k, v)
	}
}
