## Gin框架路由详解

参见：https://segmentfault.com/a/1190000016655709

Gin框架使用定制版本的`httprouter`，其路由原理是大量使用公共前缀的树结构（即前缀树`Trie tree` 或只是基数树 `Radix Tree`），具有公共前缀的节点也共享一个公共父节点。

### 01 Radix Tree

前缀树参见：https://blog.csdn.net/u013949069/article/details/78056102

基数树（Radix Tree）又称为PAT位树（Patricia Trie or crit bit tree），是一种更节省空间的前缀树（Trie Tree）。对于基数树的每个节点，如果该节点是唯一的子树的话，就和父节点合并。下图为一个基数树示例： 

 ![img](https://liwenzhou.com/images/Go/gin/radix_tree.png) 

Gin框架中注册路由的过程即为构造前缀树的过程，具有公共前缀的节点也共享一个公共父节点。如：

```go
r := gin.Default()

r.GET("/", func1)
r.GET("/search/", func2)
r.GET("/support/", func3)
r.GET("/blog/", func4)
r.GET("/blog/:post/", func5)
r.GET("/about-us/", func6)
r.GET("/about-us/team/", func7)
r.GET("/contact/", func8)
```

以上GET方法对应的路由树为：

```
Priority   Path             Handle
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>
```

其中，最右边一列每个`*<数字>`表示Handle处理函数的内存地址（一个指针），从根节点遍历到叶子节点即可得到完整的路由表。

> 与`hash-maps`不同，基数树允许通配符（路由中的参数）存在，如`/blog/:post/`中的`:post`可实现动态替换（根据路由模式进行匹配）。

***路由器为每一种请求（GET、POST、PUT、DELETE）管理一个单独的前缀树。***

为获取更好的伸缩性，每个树级别上的子节点都按照`Priority（优先级）`排序，其中的优先级就是在子节点中（子子节点...）注册的句柄的数量。优点如下：

- 优先匹配被大多数路由路径包含的节点，使尽可能多的路由被快速定位。
- 成本补偿：长路径定位花费时间长于短路径，优先匹配长路径可以达到成本补偿。

### 02 Engine结构体

#### UML结构图

 ![engine结构图](https://segmentfault.com/img/bVbh2X1?w=1374&h=774) 

#### 总结

1. `Engine`结构体内嵌了`RouterGroup`结构体，在`RouterGroup`中定义了 `GET`，`POST` 等路由注册方法。 
2. `Engine` 中的 `trees` 字段定义了路由逻辑。`trees` 是 `methodTrees` 类型（即 `[]methodTree`），`trees` 是一个数组，不同请求方法的路由在不同的树（`methodTree`）中。 
3.  `methodTree` 中的 `root` 字段（`*node`类型）是路由树的根节点。树的构造与寻址都是在 `*node`的方法中完成的。 

### 03 路由树节点

node路由树节点定义如下：

```go
type node struct {
	path      string	// 当前节点的相对路径：与所有父节点的path拼接可得到完整路径。
	indices   string    // 所有孩子节点path[0]组成的字符串：
	wildChild bool		// 子节点是否有通配符
	nType     nodeType	// 节点类型
	priority  uint32	// 优先级：当前节点以及所有子孙节点的实际路由数量。
	children  []*node 	// 子节点
    handlers  HandlersChain	// 当前节点的处理函数（包括中间件）
    fullPath  string	// 完整（原始）路径
}
```

#### path和indices

使用前缀树的逻辑：例如，两个路由分别是 `/index`，`/inter`，则根节点为 `{path: "/in", indices: "dt"...}`，两个子节点为`{path: "dex", indices: ""}，{path: "ter", indices: ""}` 。

#### handlers

`handlers`里存储了该节点对应路由下的所有处理函数，处理业务逻辑时： 

```go
func (c *Context) Next() {
    c.index++
    for s := int8(len(c.handlers)); c.index < s; c.index++ {
        c.handlers[c.index](c)
    }
}
```

除最后一个函数外，其余函数都为中间件；如果某个节点的`handler`为空，则该节点对应的路由不存在，如上述中的`/in`路由的`handlers`为空。

#### nType

Gin 中定义了四种节点类型： 

```go
const (
	static nodeType = iota 	// 普通节点，默认
	root					// 根节点
    param					// 参数路由，如/user/:id
	catchAll				// 匹配所有内容的路由，如/article/*key
)
```

`param` 与 `catchAll` 使用的区别即为 `:` 与 `*` 的区别。`*` 会把路由后面的所有内容赋值给参数 `key`；但 `:` 可以多次使用。
比如：`/user/:id/:no` 是合法的，但 `/user/*id/:no` 是非法的，因为 `*` 后面所有内容会赋值给参数 `id`。 

#### wildType

如果孩子节点是通配符（* 或 :），则该字段为true。

### 04 路由树示例

路由定义如下：

```go
r.GET("/", func(context *gin.Context) {})
r.GET("/index", func(context *gin.Context) {})
r.GET("/inter", func(context *gin.Context) {})
r.GET("/go", func(context *gin.Context) {})
r.GET("/game/:id/:k", func(context *gin.Context) {})
```

对应路由树结构图如下：

 ![路有树](https://segmentfault.com/img/bVbh22h?w=1640&h=1256) 

>最新版本中，maxParams字段转移到了Engine结构体中。

### 05 请求方法树

在gin的路由中，每一个`HTTP Method`(GET、POST、PUT、DELETE…)都对应了一棵 `radix tree`，注册路由的时候会调用下面的`addRoute`函数： 

```go
// File: gin.go

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	// ...
   
   // 获取请求方法对应的树
	root := engine.trees.get(method)
    // 无则创建
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}
```

先根据请求方法获得对应的树（每个请求方法对应一棵树），***gin使用切片保存请求方法与树的对应关系***。

> 当数据量较小时，slice查询速度优于map，原因在于：golang中的map底层使用hash实现，即需要hash函数做映射，存在函数调用开销。

`engine.trees`的类型为`methodTrees`，定义如下：

```go
type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree  // slice
```

获取请求方法对应的`get`方法定义如下：

```go
func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}
```

`engine`的初始化方法中，对`trees`字段做内存申请：

```go
func New() *Engine {
	debugPrintWARNINGNew()
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		// liwenzhou.com ...
		// 初始化容量为9的切片（HTTP1.1请求方法共9种）
		trees:                  make(methodTrees, 0, 9),
		// liwenzhou.com...
	}
	engine.RouterGroup.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}
```

### 06 注册路由

注册路由的逻辑主要有`addRoute`函数和`insertChild`方法。 

#### addRoute方法

```go
// File: tree.go

// 将具有给定句柄的节点添加到路径中。
// 非并发安全
func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++

	// 空树直接插入当前节点
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		n.nType = root
		return
	}

	parentFullPathIndex := 0

walk:
	for {
		// 找到最长的通用前缀
		// 公共前缀不包含“:”"或“*” /
		i := longestCommonPrefix(path, n.path)

		// 分割边缘
		if i < len(n.path) {
			child := node{
				path:      n.path[i:],		// 公共前缀后的部分作为子节点
				wildChild: n.wildChild,
				indices:   n.indices,
				children:  n.children,
				handlers:  n.handlers,
				priority:  n.priority - 1,	// 子节点优先级-1
				fullPath:  n.fullPath,
			}

			n.children = []*node{&child}
			// []byte for proper unicode char conversion, see #65
			n.indices = bytesconv.BytesToString([]byte{n.path[i]})
			n.path = path[:i]
			n.handlers = nil
			n.wildChild = false
			n.fullPath = fullPath[:parentFullPathIndex+i]
		}

		// 使新的节点成为当前节点的子节点
		if i < len(path) {
			path = path[i:]
			c := path[0]

			// '/' after param
			if n.nType == param && c == '/' && len(n.children) == 1 {
				parentFullPathIndex += len(n.path)
				n = n.children[0]
				n.priority++
				continue walk
			}

			// Check if a child with the next path byte exists
			for i, max := 0, len(n.indices); i < max; i++ {
				if c == n.indices[i] {
					parentFullPathIndex += len(n.path)
					i = n.incrementChildPrio(i)
					n = n.children[i]
					continue walk
				}
			}

			// Otherwise insert it
			if c != ':' && c != '*' && n.nType != catchAll {
				// []byte for proper unicode char conversion, see #65
				n.indices += bytesconv.BytesToString([]byte{c})
				child := &node{
					fullPath: fullPath,
				}
				n.addChild(child)
				n.incrementChildPrio(len(n.indices) - 1)
				n = child
			} else if n.wildChild {
				// inserting a wildcard node, need to check if it conflicts with the existing wildcard
				n = n.children[len(n.children)-1]
				n.priority++

				// 检查通配符是否匹配
				if len(path) >= len(n.path) && n.path == path[:len(n.path)] &&
					// Adding a child to a catchAll is not possible
					n.nType != catchAll &&
					// Check for longer wildcard, e.g. :name and :names
					(len(n.path) >= len(path) || path[len(n.path)] == '/') {
					continue walk
				}

				// Wildcard conflict
				pathSeg := path
				if n.nType != catchAll {
					pathSeg = strings.SplitN(pathSeg, "/", 2)[0]
				}
				prefix := fullPath[:strings.Index(fullPath, pathSeg)] + n.path
				panic("'" + pathSeg +
					"' in new path '" + fullPath +
					"' conflicts with existing wildcard '" + n.path +
					"' in existing prefix '" + prefix +
					"'")
			}

			n.insertChild(path, fullPath, handlers)
			return
		}

		// Otherwise add handle to current node
		if n.handlers != nil {
			panic("handlers are already registered for path '" + fullPath + "'")
		}
		n.handlers = handlers
		n.fullPath = fullPath
		return
	}
}
```

总结：

1. 第一次注册路由，例如注册search
2. 继续注册一条没有公共前缀的路由，例如blog
3. 注册一条与先前注册的路由有公共前缀的路由，例如support

![addroute](https://liwenzhou.com/images/Go/gin/addroute.gif) 

#### insertChild方法

```go
// File: tree.go

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {
	for {
		// Find prefix until first wildcard
		wildcard, i, valid := findWildcard(path)
		if i < 0 { // No wildcard found
			break
		}

		// The wildcard name must not contain ':' and '*'
		if !valid {
			panic("only one wildcard per path segment is allowed, has: '" +
				wildcard + "' in path '" + fullPath + "'")
		}

		// check if the wildcard has a name
		if len(wildcard) < 2 {
			panic("wildcards must be named with a non-empty name in path '" + fullPath + "'")
		}

		if wildcard[0] == ':' { // param
			if i > 0 {
				// Insert prefix before the current wildcard
				n.path = path[:i]
				path = path[i:]
			}

			child := &node{
				nType:    param,
				path:     wildcard,
				fullPath: fullPath,
			}
			n.addChild(child)
			n.wildChild = true
			n = child
			n.priority++

			// if the path doesn't end with the wildcard, then there
			// will be another non-wildcard subpath starting with '/'
			if len(wildcard) < len(path) {
				path = path[len(wildcard):]

				child := &node{
					priority: 1,
					fullPath: fullPath,
				}
				n.addChild(child)
				n = child
				continue
			}

			// Otherwise we're done. Insert the handle in the new leaf
			n.handlers = handlers
			return
		}

		// catchAll
		if i+len(wildcard) != len(path) {
			panic("catch-all routes are only allowed at the end of the path in path '" + fullPath + "'")
		}

		if len(n.path) > 0 && n.path[len(n.path)-1] == '/' {
			panic("catch-all conflicts with existing handle for the path segment root in path '" + fullPath + "'")
		}

		// currently fixed width 1 for '/'
		i--
		if path[i] != '/' {
			panic("no / before catch-all in path '" + fullPath + "'")
		}

		n.path = path[:i]

		// First node: catchAll node with empty path
		child := &node{
			wildChild: true,
			nType:     catchAll,
			fullPath:  fullPath,
		}

		n.addChild(child)
		n.indices = string('/')
		n = child
		n.priority++

		// second node: node holding the variable
		child = &node{
			path:     path[i:],
			nType:    catchAll,
			handlers: handlers,
			priority: 1,
			fullPath: fullPath,
		}
		n.children = []*node{child}

		return
	}

	// If no wildcard was found, simply insert the path and handle
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}
```

`insertChild`函数是根据`path`本身进行分割，将`/`分开的部分分别作为节点保存，形成一棵树结构。参数匹配中的`:`和`*`的区别是，前者是匹配一个字段而后者是匹配后面所有的路径。 

### 07 路由匹配

处理请求的入口函数`SeverHttp`：

```go
// File: gin.go

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // 使用对象池
	c := engine.pool.Get().(*Context)
    // Get对象后做初始化
	c.writermem.reset(w)
	c.Request = req
	c.reset()
	// 处理http请求
	engine.handleHTTPRequest(c)
	// 放回对象池
	engine.pool.Put(c)
}
```

处理http请求：

```go
// File: gin.go

func (engine *Engine) handleHTTPRequest(c *Context) {

	// 根据请求方法找到对应的路由树
	t := engine.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// 在路由树中根据path查找
		value := root.getValue(rPath, c.Params, unescape)
		if value.handlers != nil {
			c.handlers = value.handlers
			c.Params = value.params
			c.fullPath = value.fullPath
			c.Next()  // 执行函数链条
			c.writermem.WriteHeaderNow()
			return
		}
	
	c.handlers = engine.allNoRoute
	serveError(c, http.StatusNotFound, default404Body)
}
```

> 路由匹配是由节点的 `getValue`方法实现的。`getValue`根据给定的路径(键)返回`nodeValue`值，保存注册的处理函数和匹配到的路径参数数据。 

路由处理及参数匹配

```go
// File: tree.go

type nodeValue struct {
	handlers HandlersChain
	params   Params  // []Param
	tsr      bool
	fullPath string
}

func (n *node) getValue(path string, po Params, unescape bool) (value nodeValue) {
	value.params = po
walk: // Outer loop for walking the tree
	for {
		prefix := n.path
		if path == prefix {
			// 我们应该已经到达包含处理函数的节点。
			// 检查该节点是否注册有处理函数
			if value.handlers = n.handlers; value.handlers != nil {
				value.fullPath = n.fullPath
				return
			}

			if path == "/" && n.wildChild && n.nType != root {
				value.tsr = true
				return
			}

			// 没有找到处理函数 检查这个路径末尾+/ 是否存在注册函数
			indices := n.indices
			for i, max := 0, len(indices); i < max; i++ {
				if indices[i] == '/' {
					n = n.children[i]
					value.tsr = (len(n.path) == 1 && n.handlers != nil) ||
						(n.nType == catchAll && n.children[0].handlers != nil)
					return
				}
			}

			return
		}

		if len(path) > len(prefix) && path[:len(prefix)] == prefix {
			path = path[len(prefix):]
			// 如果该节点没有通配符(param或catchAll)子节点
			// 我们可以继续查找下一个子节点
			if !n.wildChild {
				c := path[0]
				indices := n.indices
				for i, max := 0, len(indices); i < max; i++ {
					if c == indices[i] {
						n = n.children[i] // 遍历树
						continue walk
					}
				}

				// 没找到
				// 如果存在一个相同的URL但没有末尾/的叶子节点
				// 我们可以建议重定向到那里
				value.tsr = path == "/" && n.handlers != nil
				return
			}

			// 根据节点类型处理通配符子节点
			n = n.children[0]
			switch n.nType {
			case param:
				// find param end (either '/' or path end)
				end := 0
				for end < len(path) && path[end] != '/' {
					end++
				}

				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[1:]
				val := path[:end]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(val); err != nil {
						value.params[i].Value = val // fallback, in case of error
					}
				} else {
					value.params[i].Value = val
				}

				// 继续向下查询
				if end < len(path) {
					if len(n.children) > 0 {
						path = path[end:]
						n = n.children[0]
						continue walk
					}

					// ... but we can't
					value.tsr = len(path) == end+1
					return
				}

				if value.handlers = n.handlers; value.handlers != nil {
					value.fullPath = n.fullPath
					return
				}
				if len(n.children) == 1 {
					// 没有找到处理函数. 检查此路径末尾加/的路由是否存在注册函数
					// 用于 TSR 推荐
					n = n.children[0]
					value.tsr = n.path == "/" && n.handlers != nil
				}
				return

			case catchAll:
				// 保存通配符的值
				if cap(value.params) < int(n.maxParams) {
					value.params = make(Params, 0, n.maxParams)
				}
				i := len(value.params)
				value.params = value.params[:i+1] // 在预先分配的容量内扩展slice
				value.params[i].Key = n.path[2:]
				if unescape {
					var err error
					if value.params[i].Value, err = url.QueryUnescape(path); err != nil {
						value.params[i].Value = path // fallback, in case of error
					}
				} else {
					value.params[i].Value = path
				}

				value.handlers = n.handlers
				value.fullPath = n.fullPath
				return

			default:
				panic("invalid node type")
			}
		}

		// 找不到，如果存在一个在当前路径最后添加/的路由
		// 我们会建议重定向到那里
		value.tsr = (path == "/") ||
			(len(prefix) == len(path)+1 && prefix[len(path)] == '/' &&
				path == prefix[:len(prefix)-1] && n.handlers != nil)
		return
	}
}
```

