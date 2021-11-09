# Zap日志库

## 一、默认的Go logger

Go语言提供默认日质包：https://golang.org/pkg/log/

### 1.1 设置Logger

设置日志记录器

```go
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Users/q1mi/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
```

### 1.2 使用Logger

示例：建立一个到URL的HTTP连接，并将状态码/错误码记录到日志文件中。

```go
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
```

### 1.3 Logger使用

```go
func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
```

执行后，生成`test.log`日志文件，并记录日志。

### 1.4 Logger特点

优点：使用简单。

缺点：仅限日志的基本级别（print）；错误日志将panic或退出程序；缺乏日志格式化；不提供日志切割功能。

## 二、Zap日志库

### 2.1 Zap简介

地址：https://github.com/uber-go/zap

优点：同时提供结构化日志记录和`printf`风格的日志记录。

### 2.2 Zap安装

安装：

```bash
$ go get -u go.uber.org/zap
```

### 2.3 配置Zap

Zap提供了两种类型的日志记录器—`Sugared Logger`和`Logger`。

在性能很好但不是很关键的上下文中，使用`SugaredLogger`。它比其他结构化日志记录包快4-10倍，并且支持结构化和printf风格的日志记录。

在每一微秒和每一次内存分配都很重要的上下文中，使用`Logger`。它甚至比`SugaredLogger`更快，内存分配次数也更少，但它只支持强类型的结构化日志记录。

#### 2.3.1 Logger

通过调用`zap.NewProduction()`/`zap.NewDevelopment()`或者`zap.Example()`创建一个Logger。

示例：

```go
var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()

	simpleHttpGet("www.github.com")
	simpleHttpGet("https://www.github.com")
	simpleHttpGet("https://www.baidu.com")

}
```

#### 2.3.2 Sugared Logger

通过`logger.Sugar()`方法获取一个`SugaredLogger`；使用`SugaredLogger`以`printf`格式记录日志。

示例：

```go
// 全局日志器
var sugarLogger *zap.SugaredLogger

// InitLogger 初始化日志器
func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugarLogger.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	// 程序关闭前刷入所有日志
	defer sugarLogger.Sync()

	simpleHttpGet("www.github.com")
	simpleHttpGet("https://www.github.com")
	simpleHttpGet("https://www.baidu.com")

}
```

### 2.4 定制Zap

#### 2.4.1 日志写入文件

使用 `zap.New(…)`方法来手动传递所有配置：

```go
func New(core zapcore.Core, options ...Option) *Logger
```

其中`zapcore.Core`需要三个配置：`Encoder`、`WriteSyncer`、`LogLevel`。

**Encoder编码器**

使用`NewJSONEncoder()`与` ProductionEncoderConfig() `。

```go
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
```

**WriteSyncer**

使用`zapcore.AddSync()`并传入打开的文件句柄。

```go
file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
//file, _ := os.Create("./test.log")
writeSyncer := zapcore.AddSync(file)
```

**Log Level**

日志级别。

示例：以JSON格式输出到文件中。

```go
// 全局日志器
var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

// InitLogger 初始化日志器
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}
```

#### 2.4.2 更改编码格式

将JSON Encoder更改为普通的Encoder：`NewJSONEncoder()`=>`Log Encoder`。

```go
func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}
```

#### 2.4.5 更改时间编码

修改时间编码器，使用大写字母记录日志级别：

```go
func getEncoder() zapcore.Encoder {
	// 修改时间编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
```

添加调用信息：

```go
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 添加调用信息
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
```

## 三、使用Lumberjack进行日志切割归档

地址：https://github.com/natefinch/lumberjack

### 3.1 安装

指令：

```bash
$ go get -u github.com/natefinch/lumberjack
```

### 3.2 使用

修改`WriteSyncer`，即修改`getLogWriter()`函数：

```go
// 日志分割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    // 最大容量
		MaxBackups: 5,     // 最大数量
		MaxAge:     30,    // 最长保存时间（天）
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
```

### 3.3 功能测试

```go
package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

// 全局日志器
var sugarLogger *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// 修改时间编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志分割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,    // 最大容量
		MaxBackups: 5,     // 最大数量
		MaxAge:     30,    // 最长保存时间（天）
		Compress:   false, // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// InitLogger 初始化日志器
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 添加调用信息
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	// 程序关闭前刷入所有日志
	defer sugarLogger.Sync()

	simpleHttpGet("https://www.github.com")
	simpleHttpGet("https://www.baidu.com")

}
```

## 四、Gin框架中使用Zap

### 4.1 Gin框架默认的中间件

Gin框架默认使用`Logger()`和`Recory()`中间件，其中，`Logger()`将gin框架本身的日志输出到标准输出，而`Recory()`在程序出现`panic`时恢复现场。

### 4.2 基于Zap的中间件

实现：

```go
// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
```

> 或可使用已封装完成的：https://github.com/gin-contrib/zap

### 4.3 在Gin中使用Zap

定义日志配置：

```go
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}
```

从配置文件加载配置信息，调用` logger.InitLogger(config.Conf.LogConfig) `完成logger实例的初始化。通过` logger.InitLogger(config.Conf.LogConfig) `注册中间件接收gin框架中的日志；使用`zap.L().Xxx()`方法记录自定义日志信息。

```go
package main

import (
	"fmt"
	"gin_zap_demo/config"
	"gin_zap_demo/logger"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	// load config from config.json
	if len(os.Args) < 1 {
		return
	}

	if err := config.Init(os.Args[1]); err != nil {
		panic(err)
	}
	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	gin.SetMode(config.Conf.Mode)

	r := gin.Default()
	// 注册zap相关中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/hello", func(c *gin.Context) {
		// 需要记录到日志的数据
		var (
			name = "lettredamour"
			age  = 24
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello https://github.com/lettredamour/!")
	})

	addr := fmt.Sprintf(":%v", config.Conf.Port)
	r.Run(addr)
}
```

