# string

## 编码方式

***Go语言使用UTF-8作为编码方式。***

- 使用unicode字符集。
- 变长编码：根据当前字符unicode码的范围分配字节个数，且每个字节的高位都为标识。

例：

| 编号          | 编码模板                   | 说明                                                 |
| ------------- | -------------------------- | ---------------------------------------------------- |
| [0, 127]      | 0???????                   | 单字节，最高位标识为0                                |
| [128, 2047]   | 110????? 10??????          | 双字节，首个字节高位标识为110，其余字节高位标识为10  |
| [2048, 65535] | 1110???? 10?????? 10?????? | 三字节，首个字节高位标识为1110，其余字节高位标识为10 |
| [65536, ...]  | ....                       | ....                                                 |

以字符`e`为例，其unicode码为101，对应的编码为`01100101`。

## 构成原理

string的构成：

- data：指向底层数组的指针（8个字节）。
- len：字节个数，标识字符串长度（8个字节）。

> string数据类型的底层结构是StringHeader，包含两个字段Data和Len，前者为指针，后者为int，各占8个字节，故string的大小为16个字节。

## 只读性

Go语言将字符串分配至只读内存段，不允许修改字符串。

### 如何修改字符串

- 装换为切片（为切片分配内存并拷贝原字符串的内容到切片内存），再通过下标修改。

- unsafe包：待续。

# slice

## 01 构成原理

slice由三部分组成：

- data：底层数组
- len：长度
- cap：容量

### 声明整型slice变量

```go
var ints []int
```

此时：

- data：nil
- len：0
- cap：0

### make整型slice变量

```go
var ints []int make([]int, 2, 5)
```

此时：

- data：指向长度为5的底层数组。
- len：当前长度（元素个数），访问超出当前长度的元素时引发panic。
- cap：

### new整型slice变量

```go
ps := new([]string)
```

此时：

- data：nil，无底层数组。
- len：0
- cap：0

通过append可为其分配内存，创建底层数组。

```go
*ps := append(*ps, "eggo")
```

字符串切片分配内存后：

- ps
  - data：
    - stirng
      - str-data：'eggo'
      - str-len：4
  - len：1
  - cap：1

## 02 底层数组

***切片并非指向底层数组的开头。***

```go
arr := [10]int{1,2,3,4,5,6,7,8,9}
var s1 []int = arr[1:4]
var s2 []int = arr[7:]
```

两个切片取自于同一个数组，将共享底层数组。

## 03 扩容规则

### 3.1 预估容量

***扩容规则***

如果原始容量的两倍小于扩容后的最小容量，则新的容量等于扩容后的最小容量。否则判断原始长度（元素个数）是否小于1024，小于则新的容量为原始容量的两倍，大于则新的容量为原始容量的1.25倍。

```go
ints := []int{1,2}
ints = append(ints, 3, 4, 5)
```

以上`oldCap = 2`，`cap = 5`，如果：

- oldCap * 2 < cap：newCap = cap
- oldCap * 2 >= cap：
  - oldLen < 1024：newCap = oldCap * 2
  - oldLen >= 1024：newCap = oldCap * 1.25

### 3.2 计算需要的内存大小

所需内存大小 = 预估容量 * 元素类型大小

### 3.3 分配内存

向内存管理模块申请内存，内存管理模块所拥有的的内存规则分别为8，16，32，48，64，80，...，内存管理模块为其分配匹配到的内存规格最合适的内存。

> 如果申请40字节内存，则为其分配48字节。

### 示例

```go
a := []string{"My", "name", "is"}
a = append(a, "eggo")
```

#### 分析

**第一步**

原始容量oldCap：3

append后至少需要的容量cap：4

判断：oldCap * 2 > cap 并且 oldCap < 1024

故新的容量newCap：oldCap * 2

# 结构体和内存对齐

示例：

```go
type T1 struct {
    a int		// 对齐值：1byte
    b int64		// 对齐值：8byte
    c int32		// 对齐值：4byte
    d int16		// 对齐值：2byte
}	// 结构体内存对齐值取字段中的最大值：8byte
```

内存占用字节必须是8的整数倍；以上结构体占用内存为24字节。

修改结构体：

```go
type T2 struct {
    a int		// 对齐值：1byte
    b int16		// 对齐值：2byte
    c int32		// 对齐值：4byte
    d int64		// 对齐值：8byte
}	// 结构体内存对齐值取字段中的最大值：8byte
```

修改后T2内存占用为16字节。

# map

## 哈希冲突处理

对于一个键值对，

1. 对键处理获得哈希值，
2. 根据hash值取对应的桶[0, m-1]
   - 取模法：hash % m
   - 与运算：hash & (m-1)
3. 哈希冲突
   - 开放地址法：桶被占用后，存储到下一个桶，查找时先找到hash对应的桶，判断键不相同时向后面的桶遍历，直至找到或遇到空桶。
   - 拉链法（常用）：桶被占用后，为当前桶添加一个链表，存储到链表上的节点。

## 哈希扩容

扩容目的：避免性能降低。

### 翻倍扩容

负载因子：Load Factor = 键值对个数count / 桶的个数 m

Go默认扩载因子（最大负载因子）为6.5。

***触发条件：负载因子超过扩载因子。***

### 等量扩容

***触发条件：负载因子未超标，但溢出桶过多。（大量键值对被删除的情况）***

- 桶的数量 <= 2^15，溢出桶的数量 >= 桶的数量的2倍。
- 桶的数量 > 2^15，溢出桶的数量 > 桶的数量。

# 函数调用栈

栈帧大小一次性分配（编译时即可确定），避免运行时越界。

栈帧布局：

- 局部变量
- 返回值
- 形参

**分析**

示例1：

```go
func inc(a int) int {
    var b int
    
    defer func() {
        a++
        b++
    }()
    
    a++
    b = a
    return b	// 将b拷贝到返回地址，defer无法修改返回地址中的值
}

func main() {
    var a, b
    b = inc(a)
    fmt.Println(a, b)	// 0, 1
}
```

示例2：

```go
func inc(a int) (b int) {
    var b int
    
    defer func() {
        a++
        b++
    }()
    
    a++
    return a	// 将a拷贝到b，执行defer时对b进行操作。
}

func main() {
    var a, b
    b = inc(a)
    fmt.Println(a, b)	// 0, 2
}
```

# 闭包

闭包的两个前提：

- 引用了外部变量。
- 即使脱离上下文，也能被调用。

# 方法

# defer

每个goroutine都存储一个链表头，新的defer添加到defer链表的头部。执行时从头部开始，故defer表现为倒序执行。

## defer节点

defer结构体

```go
type _defer struct {
    siz int32			// 参数和返回值所占字节
    started bool		// 是否已执行标志
    heap bool
    sp uintptr			// 调用者栈指针
    pc uintptr			// 返回地址
    fn *funcval			// 注册的函数
    _panic *_panic
    link *_defer		// 下一个defer节点
}
```

## 执行顺序

return最先执行并将结果写入返回值；接着执行defer；最后函数携带当前返回值退出。

## 示例

```go
func A() {
    a, b := 1, 2
    defer func(b int) {
        a = a+b
        fmt.Println(a, b)
    }(b)
    
    a= a+b
    fmt.Println(a, b)
}
```

# panic和recover

