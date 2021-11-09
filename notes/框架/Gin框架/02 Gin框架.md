## Gin框架

### Gin框架简介

地址：https://github.com/gin-gonic/gin

中文文档：https://gin-gonic.com/zh-cn/docs/

### Gin安装与使用

**安装**

```bash
$ go get -u github.com/gin-gonic/gin
```

**示例**

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context)  {
	// 200：状态码。
	// gin.H：map结构，key为string，value为interface。
	c.JSON(200, gin.H{
		"message":"Hello Golang!",
	})
}

func main() {
	// 使用默认路由引擎
	r := gin.Default()
	// 指定用户使用GET请求访问/hello时，执行sayHello函数。
	r.GET("/hello", sayHello)

	// 启动服务：指定端口，默认8080。
	err := r.Run(":9090")
	if err != nil {
		return 
	}
}
```

### RESTful API

资料：http://www.ruanyifeng.com/blog/2011/09/restful.html

REST即客户端与Web服务器之间进行交互时，使用HTTP协议中的4个请求方法代表不同的动作。

| 请求   | 动作     |
| ------ | -------- |
| GET    | 获取资源 |
| POST   | 新建资源 |
| PUT    | 更新资源 |
| DELETE | 删除资源 |

开发RESTful时使用postman测试。

**示例**

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":"DELETE",
		})
	})

	err := r.Run(":9090")
	if err != nil {
		return 
	}
}
```

