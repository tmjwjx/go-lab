package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 自定义时间格式编码器
	encoderConfig := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 日志级别大写
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05")) // 自定义时间格式
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间间隔以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 简短的调用者信息
	})

	// 创建日志核心
	core := zapcore.NewCore(
		encoderConfig,
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel, // 设置日志级别为 Debug
	)

	// 创建 Logger
	logger := zap.New(core)

	// 使用 Logger 记录日志
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Error("This is an error message")

	// 覆盖全局 Logger
	zap.ReplaceGlobals(logger)

	zap.L().Info("This is an info message", zap.String("key", "value"))
	zap.L().Debug("This is a debug message")
	zap.L().Error("This is an error message")

}

// 自定义 WriteSyncer，将日志输出到标准输出
type gologWriter struct{}


func (gologWriter) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (gologWriter) Sync() error {
	return nil
}

//package main // 定义此文件所属的包为main包，表示这是一个可执行程序的入口
//
//import (
//	"go.uber.org/zap"         // 导入Uber开发的zap日志库，提供高性能结构化日志功能
//	"go.uber.org/zap/zapcore" // 导入zap核心库，用于配置日志编码器和核心组件
//	"time"                    // 导入时间包，用于处理时间相关操作
//)
//
//func main() {
//	logger := CustomLogFormatter() // 调用自定义函数创建一个定制化的日志记录器
//	defer logger.Sync()            // 延迟调用Sync方法，确保在程序退出前将缓冲的日志刷新到输出
//
//	// 记录一条Info级别的日志，包含多个结构化字段
//	logger.Info("网络请求已接收", // 主要日志消息
//		zap.String("method", "GET"),                  // 添加字符串类型字段，记录HTTP方法
//		zap.String("path", "/api/users"),             // 添加字符串类型字段，记录请求路径
//		zap.Int("status", 200),                       // 添加整数类型字段，记录HTTP状态码
//		zap.Duration("latency", time.Millisecond*98)) // 添加持续时间类型字段，记录请求延迟
//}
//
//// CustomLogFormatter 创建并返回一个自定义配置的zap日志记录器
//func CustomLogFormatter() *zap.Logger {
//	// 自定义格式配置
//	cfg := zap.Config{
//		Level: zap.NewAtomicLevelAt(zap.InfoLevel), // 设置日志级别为Info，低于此级别的日志将不会被记录
//		//Development: false,                               // 关闭开发模式，在生产环境中使用
//		//Sampling: &zap.SamplingConfig{ // 配置日志采样策略
//		//	Initial:    100, // 在相同级别相同消息的前100条日志都会被记录
//		//	Thereafter: 100, // 之后每100条记录一次，用于减少高频日志
//		//},
//		Encoding: "console", // 使用console编码器而非json，决定了日志的输出格式
//		EncoderConfig: zapcore.EncoderConfig{ // 编码器详细配置
//			TimeKey:        "T",                                                      // 时间字段的键名为"T"
//			LevelKey:       "L",                                                      // 日志级别字段的键名为"L"
//			NameKey:        "N",                                                      // 日志记录器名称字段的键名为"N"
//			CallerKey:      "C",                                                      // 调用者信息字段的键名为"C"
//			FunctionKey:    "func",                                                   // 忽略函数名字段
//			MessageKey:     "M",                                                      // 消息内容字段的键名为"M"
//			StacktraceKey:  "S",                                                      // 堆栈跟踪字段的键名为"S"
//			LineEnding:     zapcore.DefaultLineEnding,                                // 使用默认的行结束符
//			EncodeLevel:    zapcore.CapitalColorLevelEncoder,                         // 使用大写且带颜色的日志级别编码器
//			EncodeTime:     zapcore.TimeEncoderOfLayout("[2006-01-02 15:04:05.000]"), // 自定义时间格式
//			EncodeDuration: zapcore.StringDurationEncoder,                            // 持续时间使用字符串编码
//			EncodeCaller:   zapcore.ShortCallerEncoder,                               // 使用短格式调用者信息编码器
//		},
//		OutputPaths:      []string{"stdout"}, // 日志输出到标准输出
//		ErrorOutputPaths: []string{"stderr"}, // 错误日志输出到标准错误
//	}
//
//	logger, _ := cfg.Build() // 根据配置创建日志记录器，忽略可能的错误
//	return logger            // 返回创建的日志记录器
//}
