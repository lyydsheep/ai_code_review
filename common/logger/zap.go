package log

import (
	"github.com/faiz/llm-code-review/common/enum"
	"github.com/faiz/llm-code-review/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// TODO 简化包装操作
// 使用 zap 作为日志库

var zapLogger *zap.Logger

func init() {
	InitLogger()
}

func InitLogger() {
	// 创建一个适用于生产环境的编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()

	// 设置时间编码方式为ISO8601格式，以提高日志的可读性和国际化
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 基于上述配置创建一个JSON编码器，用于生成JSON格式的日志
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	fileWriteSyncer, err := getFileLogWriter()
	if err != nil {
		panic(err)
	}
	var cores []zapcore.Core

	switch config.App.Env {
	// 只输出到文件
	case enum.ModeTEST, enum.ModePROD:
		cores = append(cores, zapcore.NewCore(encoder, fileWriteSyncer, zap.InfoLevel))

	// 同时输出到控制台和文件
	case enum.ModeDEV:
		cores = append(cores,
			zapcore.NewCore(encoder,
				zapcore.AddSync(zapcore.Lock(
					zapcore.AddSync(
						zapcore.NewMultiWriteSyncer(
							fileWriteSyncer, zapcore.AddSync(os.Stdout))))), zap.DebugLevel))

		// 需要确保 fileWriter 线程安全
		//cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		//	zapcore.NewCore(encoder, fileWriteSyncer, zap.DebugLevel))
	}
	// NewTee函数将多个core合并为一个，这样可以将日志输出到多个目的地
	core := zapcore.NewTee(cores...)

	// 创建一个新的zap.Logger，除了传入合并后的core，还添加了调用者信息和跳过一层调用栈
	// AddCaller()用于在日志中添加调用者的文件和行号信息
	// AddCallerSkip(1)用于跳过直接调用者的调用栈信息，这里设置为1，表示跳过一层
	zapLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// getFileLogWriter 返回一个同步日志写入器到文件系统。
// 该函数使用 lumberjack 作为日志滚动的实现，根据配置信息初始化 logger，
// 并通过 zapcore.AddSync 转换为 zap 的 WriteSyncer 接口。
// 配置信息包括日志文件路径、最大文件大小、日志文件保留的最大天数等。
// 返回值为 zapcore.WriteSyncer 接口，用于与 zap 日志库协同工作，
// 以及一个错误值，此函数实现中不会返回错误。
func getFileLogWriter() (zapcore.WriteSyncer, error) {
	// 初始化 lumberjack.Logger 实例，配置日志文件的路径、最大大小、最大年龄，
	// 不启用压缩，使用本地时间。
	lumberJackLogger := &lumberjack.Logger{
		Filename:  config.App.Log.Path,
		MaxSize:   config.App.Log.MaxSize,
		MaxAge:    config.App.Log.MaxAge,
		Compress:  false,
		LocalTime: true,
	}
	// 使用 zapcore.AddSync 将 lumberjack.Logger 转换为 zapcore.WriteSyncer 接口，
	// 以便在 zap 日志库中使用，并返回。
	return zapcore.AddSync(lumberJackLogger), nil
}

// TODO 暂时用于测试
func ZapLoggerTest() {
	zapLogger.Info("Hello, World!", zap.Any("key", "value"))
}
