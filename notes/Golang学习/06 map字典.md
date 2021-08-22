## map定义

- 方式一

  ```go
  // 声明
  var m1 map[string]string
  // 初始化
  m1 = map[string]string{}
  // 赋值
  m1["name"] = "ohiua"
  m1["sex"] = "man"
  fmt.Println(m1)
  ```

- 方式二

  ```go
  // 方式二
  m2 := map[string]string{}
  m2["name"] = "ohiua"
  m2["sex"] = "man"
  fmt.Println(m2)
  ```

- 方式三

  ```go
  // 方式三
  m3 := make(map[string]string)
  m3["name"] = "ohiua"
  m3["sex"] = "man"
  fmt.Println(m3)
  ```

## map赋值

**map的value可存入任意类型：**

```go
func main() {

	m := map[int]interface{}{}
	m[1] = 1
	m[2] = false
	m[3] = "str"
	m[4] = []int{1,2,3}

	fmt.Println(m)
}
```

## map删除元素

```go
func main() {

	m := map[int]interface{}{}
	m[1] = 1
	m[2] = false
	m[3] = "str"
	m[4] = []int{1,2,3}

	fmt.Println("m length: " , len(m))
    delete(m, 4)	// key: 4
	fmt.Println("m length: " , len(m))
}
```

## map遍历

```go
func main() {

	m := map[string]interface{}{}
	m["a"] = 1
	m["b"] = false
	m["c"] = "str"
	m["d"] = []int{1,2,3}

	for k, v := range m {
		fmt.Printf("m[%s] = %v\n", k, v)
	}
}
```

