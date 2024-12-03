package Tlog

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	Level             string
	DisableCaller     bool
	DisableStacktrace bool
	OutputPaths       []string
	format            string
	ErrorOutputPaths  []string
}

func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths:       []string{"stdout", "./tiktokk.log"},
		format:            "console",
		ErrorOutputPaths:  []string{"stderr"},
	}
}

func LogOption() *Options {
	return &Options{
		Level:             viper.GetString("log.level"),
		DisableCaller:     viper.GetBool("log.disableCaller"),
		DisableStacktrace: viper.GetBool("log.disableStackTrace"),
		format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.outputPaths"),
		ErrorOutputPaths:  viper.GetStringSlice("log.errorOutputPaths"),
	}
}
