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
