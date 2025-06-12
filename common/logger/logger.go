package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

// logger 符合链式调用
type logger struct {
	ctx      context.Context
	traceId  string
	spanId   string
	parentId string
	logger   *zap.Logger
}

func New(ctx context.Context) *logger {
	var traceId, spanId, parentId string
	if ctx.Value("traceId") != nil {
		traceId = ctx.Value("traceId").(string)
	}
	if ctx.Value("spanId") != nil {
		spanId = ctx.Value("spanId").(string)
	}
	if ctx.Value("parentId") != nil {
		parentId = ctx.Value("parentId").(string)
	}
	return &logger{
		ctx:      ctx,
		traceId:  traceId,
		spanId:   spanId,
		parentId: parentId,
		logger:   zapLogger,
	}
}

func (l *logger) Debug(msg string, kv ...any) {
	l.log(zapcore.DebugLevel, msg, kv...)
}

func (l *logger) Info(msg string, kv ...any) {
	l.log(zapcore.InfoLevel, msg, kv...)
}

func (l *logger) Warn(msg string, kv ...any) {
	l.log(zapcore.WarnLevel, msg, kv...)
}

func (l *logger) Error(msg string, kv ...any) {
	l.log(zapcore.ErrorLevel, msg, kv...)
}

func (l *logger) Fatal(msg string, kv ...any) {
	l.log(zapcore.FatalLevel, msg, kv...)
}

func (l *logger) log(level zapcore.Level, msg string, kv ...any) {
	if ce := l.logger.Check(level, msg); ce != nil {
		if len(kv) == 0 {
			ce.Write()
			return
		}
		// 确保 kv 长度为偶数，并记录不匹配的情况
		if len(kv)%2 != 0 {
			kv = append(kv, "invalid_value")
			l.logger.Warn("Odd number of key-value pairs provided, appending 'invalid_value'", zap.Int("length", len(kv)))
		}
		kv = append(kv, "traceId", l.traceId, "spanId", l.spanId, "parentId", l.parentId)
		funcName, file, line := getCallerInfo()
		kv = append(kv, "FuncName", funcName, "File", file, "Line", line)
		fields := make([]zap.Field, len(kv)/2)
		for i := 0; i < len(kv)/2; i++ {
			fields[i] = zap.Any(fmt.Sprintf("%v", kv[i*2]), kv[i*2+1])
		}
		ce.Write(fields...)
	} else {
		l.logger.Warn("Failed to check log level or message", zap.String("level", level.String()), zap.String("msg", msg))
	}
}

func getCallerInfo() (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return "", "", 0
	}
	file = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return funcName, file, line
}
