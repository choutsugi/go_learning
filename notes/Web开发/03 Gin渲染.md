## Gin渲染

前后端不分离的Web架构中，在后端将数据渲染到HTML文档中，从而实现动态的网页效果。模板即事先定义好的HTML文档，模板渲染机制类似于文本替换，使用相应的数据替换掉HTML文档中的标记。

### HTML渲染

创建文件夹`templates`存放模板文件，按业务分别创建子文件夹`posts`和`user`。

`post/index.html`文件内容如下：

```html
{{define "posts/index.html"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}
```

`users/index.html`文件内容如下：

```html
{{define "users/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>users/index</title>
</head>
<body>
{{.title}}
</body>
</html>
{{end}}
```

Gin框架中使用`LoadHTMLGlob()`或者`LoadHTMLFiles()`方法进行HTML模板渲染：

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
```

### 自定义模板函数

 定义一个不转义相应内容的`safe`模板函数如下： 

```go
package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://bilibili.com'>Bilibili</a>")
	})

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
```

 在`index.tmpl`中使用定义好的`safe`模板函数： 

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>修改模板引擎的标识符</title>
</head>
<body>
<div>{{ . | safe }}</div>
</body>
</html>
```

### 静态文件处理

当渲染的HTML文件中引用了静态文件时，只需要按照以下方式在渲染页面前调用`gin.Static`方法即可。 

```go
func main() {

	r := gin.Default()
	// 引用静态文件
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
```

### 使用模板继承

Gin框架默认都是使用单模板，如果需要使用`block template`功能，可以通过`"github.com/gin-contrib/multitemplate"`库实现，具体示例如下：

项目目录下的templates文件夹下有以下模板文件，其中`home.tmpl`和`index.tmpl`继承了`base.tmpl`：

```
templates
├── includes
│   ├── home.tmpl
│   └── index.tmpl
├── layouts
│   └── base.tmpl
└── scripts.tmpl
```

定义`loadTemplates`函数：

```go
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
```

main函数：

```go
func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.GET("/index", indexFunc)
	r.GET("/home", homeFunc)

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
```

### 补充文件路径处理

 模板文件和静态文件的路径根据项目设置，可使用以下函数获取当前执行程序的路径：

```go
package utils

import (
	"os"
	"path/filepath"
)

func getCurrentPath() string {
	executable, err := os.Executable()
	if err != nil {
		return "./"
	}
	return filepath.Dir(executable)
}
```

### Json渲染

```go
func main() {
	r := gin.Default()

	r.GET("/someJson", func(context *gin.Context) {
		// 方式一：自行拼接Json
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello Golang!",
		})
	})

	r.GET("/moreJson", func(context *gin.Context) {
		// 方式二：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}

		msg.Name = "stone"
		msg.Message = "Hello Golang!"
		msg.Age = 18
		context.JSON(http.StatusOK, msg)
	})

	err := r.Run(":60000")
	if err != nil {
		return
	}
}
```

### XML渲染

```go
// XML渲染
r.GET("/someXML", func(context *gin.Context) {
    // 方式一：自行拼接
    context.XML(http.StatusOK, gin.H{
        "message": "Hello Golang!",
    })
})

r.GET("/moreXML", func(context *gin.Context) {
    // 方式二：使用结构体
    type MessageRecord struct {
        Name    string
        Message string
        Age     int
    }

    var msg MessageRecord
    msg.Name = "stone"
    msg.Message = "Hello Golang!"
    msg.Age = 18
    context.XML(http.StatusOK, msg)
})
```

### Yaml渲染

```go
// YAML渲染
r.GET("someYAML", func(context *gin.Context) {
    context.YAML(http.StatusOK, gin.H{
        "message": "ok",
        "status":  http.StatusOK,
    })
})
```

### protobuf渲染

```go
// protobuf渲染
r.GET("someProtobuf", func(context *gin.Context) {
    reps := []int64{int64(1), int64(2)}
    label := "test"
    data := &protoexample.Test{
        Label: &label,
        Reps:  reps,
    }
    // 输出被序列化的二进制数据。
    context.ProtoBuf(http.StatusOK, data)
})
```









