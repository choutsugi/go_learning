## Gin环境搭建

- Go版本：大于1.11（go mod）。
- 环境变量：GOPROXY=https://goproxy.io，使用国内镜像。
- GO111MODULE = auto，默认，无需设置。
- go mod init：```go mod init GinFrame``` 或 使用 goland 创建 go mod 项目。

## Gin资料

官方地址：https://github.com/gin-gonic/gin 

中文资料：https://www.kancloud.cn/shuangdeyu/gin_book

## 安装Postman

下载地址：https://www.postman.com/downloads/

## 测试

1. 使用Goland新建GinLearning项目（Go mod）并设置GOPROXY。

2. 新建main文件：

   ```go
   package main
   
   import "github.com/gin-gonic/gin"
   
   func main() {
   	r := gin.Default()	// 携带基础中间件启动
   	r.GET("/ping", func(c *gin.Context){
   		c.JSON(200, gin.H{
   			"message":"pong",
   		})
   	})
   
   	r.Run()				// 监听localhost:8080
   }
   ```

3. 浏览器访问http://127.0.0.1:8080/ping，显示Json：```{"message":"pong"}```即成功。

