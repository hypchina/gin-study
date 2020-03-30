package test

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZapLog(t *testing.T) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		Level:         atom,          // 日志级别
		Development:   false,         // 开发模式，堆栈跟踪
		Encoding:      "json",        // 输出格式 console 或 json
		EncoderConfig: encoderConfig, // 编码器配置
		//OutputPaths:      []string{"stderr", "./logs/spikeProxy.log"},       // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		//ErrorOutputPaths: []string{"stderr"},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}
	sugar := logger.Sugar()
	sugar.Error("fuck u error")
	sugar.Debug("fuck u debug")
	sugar.Info("fuck u info")
}
