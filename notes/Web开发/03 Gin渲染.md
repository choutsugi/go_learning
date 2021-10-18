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

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
```

