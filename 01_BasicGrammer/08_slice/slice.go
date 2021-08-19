/*
	切片：动态数组。
*/

package main

import "fmt"

// 打印静态数组（值传递）
func printArray(arr [4]int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	arr[0] = -1 // 不改变实参
	fmt.Println("")
}

// 打印动态数组（引用传递）
func printDynamicArray(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	arr[0] = -1 // 改变实参
	fmt.Println("")
}

func main() {

	// 静态数组与动态数组的区别
	// 静态数组
	var arr1 [10]int
	arr2 := [10]int{1, 2, 3, 4}
	arr3 := [4]int{11, 22, 33, 44}

	fmt.Printf("arr1: ")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Printf("\n")

	for index, value := range arr2 {
		fmt.Println("arr2 index = ", index, ", value = ", value)
	}

	// 查看静态数组数据类型
	fmt.Printf("type of arr1: %T\n", arr1)
	fmt.Printf("type of arr2: %T\n", arr2)
	fmt.Printf("type of arr3: %T\n", arr3)

	// 打印静态数组
	fmt.Println("===print static array===")
	printArray(arr3)
	printArray(arr3)

	// 动态数组slice
	dynamic_arr_01 := []int{1, 2, 3, 5, 7, 11}
	fmt.Println("===print dynamic array===")
	printDynamicArray(dynamic_arr_01)
	printDynamicArray(dynamic_arr_01)
	// 查看动态数组数据类型
	fmt.Printf("type of dynamic_arr_01: %T\n", dynamic_arr_01)

	// 声明动态数组
	// 方法一：声明并初始化
	// dynamic_arr_02 := []int{1, 2, 3}
	// 方法二：仅声明（不分配空间）
	// var dynamic_arr_02 []int
	// dynamic_arr_02 = make([]int, 3) // 开辟3个空间，默认值为0
	// 方法三：声明并分配空间
	// var dynamic_arr_02 = make([]int, 3)
	// 方法四：声明并分配空间，自动推导
	dynamic_arr_02 := make([]int, 3)
	fmt.Printf("size = %d, capacity = %d, dynamic_arr_02 = %v\n", len(dynamic_arr_02), cap(dynamic_arr_02), dynamic_arr_02) // size = 3, capacity = 3, slice = [0 0 0]

	// 判断切片是否为空
	if dynamic_arr_02 == nil {
		fmt.Println("dynamic_arr_02 is empty!")
	} else {
		fmt.Println("dynamic_arr_02 isn't empty!")
	}

	// 动态数组追加元素
	var numbers = make([]int, 3, 5) // 指定大小和容量

	fmt.Printf("size = %d, capacity = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 追加元素：size = 4, capacity = 5, slice = [0 0 0 1]
	numbers = append(numbers, 1)
	fmt.Printf("size = %d, capacity = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 追加元素：size = 5, capacity = 5, slice = [0 0 0 1 2]
	numbers = append(numbers, 2)
	fmt.Printf("size = %d, capacity = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 追加元素：size = 6, capacity = 10, slice = [0 0 0 1 2 3]
	// 扩容因子：2
	numbers = append(numbers, 3)
	fmt.Printf("size = %d, capacity = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 动态数组截取元素
	dynamic_arr_03 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// 打印原始切片
	fmt.Println("dynamic_arr_03[]: ", dynamic_arr_03)
	// 打印子切片：索引0（包含）到索引4（不包含）
	fmt.Println("dynamic_arr_03[0:4]: ", dynamic_arr_03[0:4])
	// 子切片默认下限为0
	fmt.Println("dynamic_arr_03[ :4]: ", dynamic_arr_03[:4])
	// 子切片默认上限为最后一个索引
	fmt.Println("dynamic_arr_03[ 4:]: ", dynamic_arr_03[4:])

	// 子切片与切片指向同一块内存数据
	dynamic_arr_04 := []int{1, 2, 3}
	dynamic_arr_05 := dynamic_arr_04[0:2]

	dynamic_arr_04[0] = -1
	fmt.Println(dynamic_arr_04)
	fmt.Println(dynamic_arr_05)

	// 拷贝切片内容
	dynamic_arr_06 := make([]int, 3) // 申请新空间然后拷贝
	copy(dynamic_arr_06, dynamic_arr_05)
	dynamic_arr_05[0] = 1
	fmt.Println(dynamic_arr_05)
	fmt.Println(dynamic_arr_06)

}
