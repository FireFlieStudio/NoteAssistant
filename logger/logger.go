package logger

import (
	"NoteAssistant/conf"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logPath                        string
	IsDebugOn, localTime, compress bool
	maxSize, maxBackups, maxAge    int
	Logger                         *zap.SugaredLogger
)

func setDefaultLogConfig() {
	viper.SetDefault("log.debug", true)
	viper.SetDefault("log.logPath", "logs/")

	viper.SetDefault("log.maxSize", 200)
	viper.SetDefault("log.maxBackups", 10)
	viper.SetDefault("log.maxAge", 30)
	viper.SetDefault("log.localTime", true)
	viper.SetDefault("log.compress", true)
}

func loadLogConfig() {
	IsDebugOn = viper.GetBool("log.debug")
	logPath = viper.GetString("log.logPath")

	maxSize = viper.GetInt("log.maxSize")
	maxBackups = viper.GetInt("log.maxBackups")
	maxAge = viper.GetInt("log.maxAge")
	localTime = viper.GetBool("log.localTime")
	compress = viper.GetBool("log.compress")
}

func init() {
	conf.InitConfig()
	setDefaultLogConfig()
	loadLogConfig()

	Logger = NewLogger()
}

func Debug(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}
