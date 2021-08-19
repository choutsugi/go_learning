/*
	包。
*/

package main

import (
	_ "05_init/lib1"      // 包匿名别名：不使用该包，只执行其init()方法。
	mylib2 "05_init/lib2" // 包具名别名：
	// . "05_init/lib2"   // 导入原包的全部方法到当前包，无需通过原包名调用其中方法（慎用）。
)

func main() {
	// lib1.Lib1Test()
	mylib2.Lib2Test()
}
