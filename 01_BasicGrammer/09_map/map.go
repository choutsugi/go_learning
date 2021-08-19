/*
	map
*/

package main

import "fmt"

// map传参：引用传递
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Print("key = ", key)
		fmt.Print(", value =", value)
		fmt.Println()
	}
}

func main() {
	// 第一种声明方法：
	// 声明map类型：key为string，value为string。
	var mMap1 map[string]string

	if mMap1 == nil {
		fmt.Println("mMap1 is empty!")
	}

	// 分配数据空间
	mMap1 = make(map[string]string, 10)

	mMap1["teemo"] = "小可爱"
	mMap1["yasuo"] = "快乐"
	mMap1["zoe"] = "捣蛋鬼"

	// 哈希表：乱序
	fmt.Println(mMap1)

	// 第二种声明方法：
	mMap2 := make(map[int]string)
	mMap2[1] = "teemo"
	mMap2[2] = "yasuo"
	mMap2[3] = "zoe"
	fmt.Println(mMap2)

	// 第三种声明方法：
	mMap3 := map[string]string{
		"one":   "C++",
		"two":   "Python",
		"three": "Go",
	}
	fmt.Println(mMap3)

	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["UK"] = "London"
	// 删除
	delete(cityMap, "Japan")
	// 修改
	cityMap["Japan"] = "None"
	// 遍历（for-range）
	printMap(cityMap)
	// 拷贝：遍历=>重新赋值
}
