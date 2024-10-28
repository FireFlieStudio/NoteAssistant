package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

func NewLogger() *zap.SugaredLogger {
	coreList := make([]zapcore.Core, 0)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// DEBUG Stdout Logger
	if IsDebugOn {
		coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	}

	// INFO Logger
	infoJackLogger := &lumberjack.Logger{
		Filename:   path.Join(logPath, "info.log"),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(infoJackLogger), zapcore.InfoLevel))

	// Error Logger
	errJackLogger := &lumberjack.Logger{
		Filename:   path.Join(logPath, "error.log"),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
	}
	coreList = append(coreList, zapcore.NewCore(encoder, zapcore.AddSync(errJackLogger), zapcore.ErrorLevel))

	// Collect all cores
	core := zapcore.NewTee(coreList...)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
	return zapLogger.Sugar()
}
