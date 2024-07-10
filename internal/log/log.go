package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string `json:"level" yaml:"level" env:"LOG_LEVEL" default:"info"`
}

func Init(config Config) (log *zap.Logger, err error) {
	cores := make([]zapcore.Core, 0)

	// console log
	cores = append(cores, createConsoleCore())

	core := zapcore.NewTee(cores...)
	caller := zap.AddCaller()
	callerSkip := zap.AddCallerSkip(2)
	logger := zap.New(core, caller, callerSkip, zap.Development())
	zap.ReplaceGlobals(logger)
	if _, err := zap.RedirectStdLogAt(logger, zapcore.ErrorLevel); err != nil {
		return nil, err
	}

	return logger, nil
}

func createConsoleCore() zapcore.Core {
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = timeEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	return zapcore.NewCore(consoleEncoder, consoleDebugging, zapcore.DebugLevel)
}

// customCallerEncoder set caller fullpath
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.FullPath())
}

// timeEncoder format time
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
