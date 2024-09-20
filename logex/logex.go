package logex

import (
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
)

func HookLog() {
	logrus.AddHook(NewRotateHook())
}

func NewRotateHook() logrus.Hook {
	hook, _ := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   "logs/output.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     3, //days
			Compress:   false,
			LocalTime:  true,
		},
		logrus.DebugLevel,
		&logrus.TextFormatter{DisableColors: true},
		&lumberjackrus.LogFileOpts{
			logrus.TraceLevel: &lumberjackrus.LogFile{
				Filename:   "logs/trace.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.DebugLevel: &lumberjackrus.LogFile{
				Filename:   "logs/debug.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.WarnLevel: &lumberjackrus.LogFile{
				Filename:   "logs/warn.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.InfoLevel: &lumberjackrus.LogFile{
				Filename:   "logs/info.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.ErrorLevel: &lumberjackrus.LogFile{
				Filename:   "logs/error.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.FatalLevel: &lumberjackrus.LogFile{
				Filename:   "logs/fatal.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
			logrus.PanicLevel: &lumberjackrus.LogFile{
				Filename:   "logs/panic.log",
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     3, //days
				Compress:   false,
				LocalTime:  true,
			},
		},
	)
	return hook
}
