package logger

import (
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger
var level = &atomic.String{}

func init() {
	SetLevel(zapcore.DebugLevel)
	encoder := zapcore.NewConsoleEncoder(DefaultEncoderConfig())
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(DefaultConsoleSyncer())
	core := zapcore.NewCore(encoder, multiWriteSyncer, DefaultLevelEnabler())
	logger = zap.New(core, zap.AddCaller())
	_ = logger.Sync()
	sugarLogger = logger.Sugar()
	_ = sugarLogger.Sync()
}

func SetLevel(l zapcore.Level) {
	level.Store(l.String())
}

func DefaultLevelEnabler() zap.LevelEnablerFunc {
	return func(z zapcore.Level) bool {
		l, _ := zapcore.ParseLevel(level.Load())
		return z >= l
	}
}

func DefaultTimeEncoder() (timeEncoder zapcore.TimeEncoder) {
	timeEncoder = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	return
}

func DefaultEncoderConfig() (encoderConfig zapcore.EncoderConfig) {
	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = DefaultTimeEncoder()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return
}

func DefaultConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func Log() *zap.SugaredLogger {
	return sugarLogger
}
