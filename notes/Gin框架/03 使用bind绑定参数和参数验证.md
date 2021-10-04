## 模型绑定：bind模式的使用

### MustBind

使用ShouldBind。

### ShouldBind

#### 不同方式传参

**通过body传参**

用这个。

```go
// PostParams 接受前端传回的json数据。
type PostParams struct {
	Name string `json:"name"`	// tag
	Age int `json:"age"`
	Sex bool `json:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/testBind", func(c *gin.Context){
		var p PostParams
		// 绑定前端传回的json数据。
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":"Wrong!",
				"data":gin.H{},
			})
		}else{
			c.JSON(200, gin.H{
				"msg":"Success",
				"data":p,
			})
		}
	})

	router.Run(":8081") // 监听localhost:8080
}
```

**Postman：**

- 地址： localhost:8081/testBind 

- 请求方式：POST

- 请求数据：Body => raw

  ```json
  {
      "name":"ohiua",
      "age":24,
      "sex":true
  }
  ```

**通过uri传参**

```go
// PostParams 接受前端传回的json数据。
type PostParams struct {
	Name string `json:"name" uri:"name"` // tag
	Age int `json:"age" uri:"age"`
	Sex bool `json:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/testBind/:name/:age/:sex", func(c *gin.Context){
		var p PostParams
		// 绑定前端传回的uri数据。
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":"Wrong!",
				"data":gin.H{},
			})
		}else{
			c.JSON(200, gin.H{
				"msg":"Success",
				"data":p,
			})
		}
	})

	router.Run(":8081") // 监听localhost:8080
}
```

**Postman：**

- 地址： localhost:8081/testBind/ohiua/24/true 
- 请求方式：POST。

**通过url传参**

```go
// PostParams 接受前端传回的json数据。
type PostParams struct {
	Name string `json:"name" uri:"name" form:"name"` // tag
	Age int `json:"age" uri:"age" form:"age"`
	Sex bool `json:"sex" uri:"sex" form:"sex"`
}

func main() {
	router := gin.Default()

	router.POST("/testBind", func(c *gin.Context){
		var p PostParams
		// 绑定前端传回的url数据。
		err := c.ShouldBindQuery(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":"Wrong!",
				"data":gin.H{},
			})
		}else{
			c.JSON(200, gin.H{
				"msg":"Success",
				"data":p,
			})
		}
	})

	router.Run(":8081") // 监听localhost:8080
}
```

**Postman：**

- 地址： localhost:8081/testBind?name=ohiua&age=24&sex=true 
- 请求方式：POST

## 表单验证

### 默认表单验证

```go
// PostParams 接受前端传回的json数据。
type PostParams struct {
	Name string `json:"name" binding:"required"`	// 表单验证
	Age int `json:"age" binding:"required"`
	Sex bool `json:"sex" binding:"required"`
}
```

**Postman：**

- 地址： localhost:8081/testBind 

- 请求方式：POST

- 请求数据：Body => raw

  ```go
  {"age":24}
  ```

- 返回结果：表单验证不通过，请求参数中name和sex为空。

  ```go
  {
      "data": {},
      "msg": "Wrong!"
  }
  ```

控制台打印错误信息：

```
Key: 'PostParams.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'PostParams.Sex' Error:Field validation for 'Sex' failed on the 'required' tag
```

### 自定义表单验证器

```go
// PostParams 接受前端传回的json数据。
type PostParams struct {
	Name string `json:"name"` 	// tag
	Age int `json:"age" binding:"required,mustBig"`	// 使用mustBig标签校验age属性
	Sex bool `json:"sex"`
}

// 自定义表单验证规则：对age进行过滤。
func mustBig(fl validator.FieldLevel) bool{
	fmt.Println(fl.Field().Interface().(int))
	if fl.Field().Interface().(int) >= 18 {
		return true
	}
	return false
}

func main() {
	router := gin.Default()

	// 表单验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个表单验证规则
		v.RegisterValidation("mustBig", mustBig)	// 使用标签名+函数名
	}

	//...
}
```

**Postman：**

- 地址： localhost:8081/testBind 

- 请求方式：POST

- 请求数据：Body => raw

  ```json
  {"age":18}
  ```

  - 返回结果：

    ```json
    {
        "data": {
            "name": "",
            "age": 18,
            "sex": false
        },
        "msg": "Success"
    }
    ```

- 请求数据：Body => raw

  ```go
  {"age":17}
  ```

  - 返回结果：

    ```go
    {
        "data": {},
        "msg": "Wrong!"
    }
    ```

