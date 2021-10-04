## 字典

map是一种无序的基于`key-value`的数据结构，内部使用```散列表（hash）```实现。map是引用类型，默认初始值为nil，必须初始化才能使用。 使用make()函数分配内存，语法为：```make(map[KeyType]ValueType, [cap])```。

### map定义

**先声明后初始化：**

```go
// 声明map类型：未初始化，初始值为nil
var a map[string]int
fmt.Println(a == nil)
// map 初始化
a = make(map[string]int, 8)
fmt.Println(a == nil)
```

**声明的同时进行初始化：**

```go
b := map[int]bool{
    1: true,
    2: false,
}
fmt.Printf("b: %#v \ntype: %T\n", b, b)
```

### map添加数据

```go
a["OHIUA"] = 24
a["SHINRIN"] = 23
fmt.Printf("a: %#v \ntype: %T\n", a, a)
```

**map的value可存入任意类型：**

```go
m := map[int]interface{}{}
m[1] = 1
m[2] = false
m[3] = "str"
m[4] = []int{1,2,3}

fmt.Println(m)
```

### map查询元素

判断键是否存在。

```go
var scoreMap = make(map[string]int, 8)
scoreMap["YaSuo"] = 100
scoreMap["Teemo"] = 99
v, ok := scoreMap["Teemo"]
if ok {
    fmt.Println("Teemo in the scoreMap:", v)
} else {
    fmt.Println("Teemo is not in the scoreMap.")
}
```

### map遍历

```go
for key, value := range scoreMap {
    fmt.Println(key, value)
}
```

### map删除元素

```go
delete(scoreMap, "YaSuo")
fmt.Println(scoreMap)
```

### map按指定顺序遍历

```go
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
```

### 元素类型为map的切片

```go
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
```

### 元素类型为切片的map

```go
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
```

### map应用：统计字符串中每个单词出现的次数

```go
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
```

