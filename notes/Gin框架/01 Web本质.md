## Web简介

### 简介

- Web是基于HTTP协议进行交互的应用网络。
- Web是通过使用浏览器/APP访问的各种资源。

> 浏览器向服务器发送请求，服务器向浏览器发送响应。

### 示例

浏览器发送一个请求，服务端返回一个文件内容：

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	bytes, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		return
	}
	_, _ = fmt.Fprintf(w, string(bytes))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}
}
```





