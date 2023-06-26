package Tlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

type zapLogger struct {
	z *zap.Logger
}

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

var _ Logger = (*zapLogger)(nil)

var (
	std *zapLogger
	mu  sync.Mutex
)

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = NewLogger(opts)
}

func NewLogger(op *Options) *zapLogger {
	if op == nil {
		op = NewOptions()
	}
	//日志等级转化
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(op.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	proConfig := zap.NewProductionEncoderConfig()
	proConfig.MessageKey = "massage"
	proConfig.TimeKey = "timestamp"
	proConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	proConfig.EncodeDuration = func(duration time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendFloat64(float64(duration) / float64(time.Millisecond))
	}
	cfg := &zap.Config{
		DisableCaller:     op.DisableCaller,
		DisableStacktrace: op.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          op.format,
		EncoderConfig:     proConfig,
		OutputPaths:       op.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"}}
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z}
	zap.RedirectStdLog(z)
	return logger
}

func (z *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Debugw(msg, keysAndValues...)
}

func (z *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Infow(msg, keysAndValues...)
}

func (z *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Warnw(msg, keysAndValues...)
}

func (z *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Errorw(msg, keysAndValues...)
}

func (z *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Panicw(msg, keysAndValues...)
}

func (z *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	z.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (z *zapLogger) Sync() {
	z.z.Sugar().Sync()
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func Sync() {
	std.z.Sugar().Sync()
}
