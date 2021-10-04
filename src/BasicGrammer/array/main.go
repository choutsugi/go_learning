package main

import "fmt"

func changeArray(arr [3]int) {
	arr[0] = 100
}

// 数组
func main() {
	var a [3]int
	var b [4]int

	fmt.Println(a)
	fmt.Println(b)

	// 数组初始化
	// 1.定义的同时使用初始值列表来初始化
	var heroArray = [4]string{"Teemo", "YaSuo", "LiQing", "ZhaoXin"}
	fmt.Println(heroArray)
	fmt.Println(heroArray[2])
	// 2.编译器推导数组的长度
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)
	// 3.使用索引值方式初始化
	var langArray = [...]string{1: "CN", 3: "KR", 7: "JP"}
	fmt.Println(langArray)
	fmt.Printf("type: %T  len: %d\n", langArray, len(langArray)) // type: [8]string  len: 8

	// 数组遍历
	// 1.for循环遍历
	var animalArray = [4]string{"CAT", "DOG", "FISH", "BIRD"}
	for i := 0; i < len(animalArray); i++ {
		fmt.Println(animalArray[i])
	}
	// 2.for range遍历
	for _, v := range animalArray {
		fmt.Println(v)
	}

	// 二维数组
	cityArray := [...][2]string{
		{"Bei'ping", "Xi'an"},
		{"Chong'qing", "Cheng'du"},
		{"Hang'zhou", "Shang'hai"},
	}
	fmt.Println(cityArray)

	// 二维数组的遍历
	for _, v1 := range cityArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x) // 1, 2, 3
	changeArray(x)
	fmt.Println(x) // 1, 2, 3
}
