## 修改端口号

```go
gin.Default().Run(":8081")
```

## GET请求

查：参数挂在url中 或 使用uri传参。

**url传参：**

```go
// 请求地址：localhost:8081/path/123456?user=ohiua&pwd=123456

router.GET("/path/:id", func(c *gin.Context){
    // 获取url通过占位符传入的参数
    id := c.Param("id")
    // 获取uri传入的参数
    user := c.DefaultQuery("user", "OHIUA")	// 设置默认值
    pwd := c.Query("pwd")
    c.JSON(200, gin.H{
        "id":id,
        "user":user,
        "pwd":pwd,
    })
})
```

## POST请求

增：参数在 form body 或 uri 中。

```go
// 请求地址：localhost:8081/path/

router.POST("/path", func(c *gin.Context){
    user := c.DefaultPostForm("user", "OHIUA")
    pwd := c.PostForm("pwd")
    c.JSON(200, gin.H{
        "user":user,
        "pwd":pwd,
    })
})
```

## DELETE请求

删：参数在 body 或 uri 中。

```go
// 请求地址：localhost:8081/path/123

router.DELETE("/path/:id", func(c *gin.Context){
    id := c.Param("id")
    c.JSON(200, gin.H{
        "id":id,
    })
})
```

## PUT请求

改：参数在 form body 或 uri 中。

```go
// 请求地址：localhost:8081/path/

router.PUT("/path", func(c *gin.Context){
    user := c.DefaultPostForm("user", "OHIUA")
    pwd := c.PostForm("pwd")
    c.JSON(200, gin.H{
        "user":user,
        "pwd":pwd,
    })
})
```

## 如何获取参数

- GET：一般情况下为地址栏的Query。
- POST、DELETE、PUT：一般为form参数。
- URL：地址栏定义占位符，通过占位符取参数。

## restFul规范

待续。

