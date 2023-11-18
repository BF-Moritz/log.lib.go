package loglibgo

import (
	"fmt"
	"os"
	"runtime"

	"github.com/BF-Moritz/log.lib.go/consts"
	"github.com/BF-Moritz/log.lib.go/enum"
)

type Logger struct {
	level            enum.LogLevel
	showTime         bool
	timeFormatString string

	file *os.File

	colorReset string
	colorDebug string
	colorInfo  string
	colorError string
	colorFatal string
}

// NewLogger Creates a new Logger
func NewLogger(level enum.LogLevel, showTime, showColor bool, file *os.File) *Logger {
	var logger Logger = Logger{
		level:            level,
		showTime:         showTime,
		timeFormatString: consts.DefaultTimeFormatString,

		file: file,

		colorReset: consts.ColorReset,
		colorDebug: consts.ColorCyan,
		colorInfo:  consts.ColorGreen,
		colorError: consts.ColorRed,
		colorFatal: consts.ColorPurple,
	}

	if runtime.GOOS == "windows" || !showColor {
		logger.colorReset = ""
		logger.colorDebug = ""
		logger.colorInfo = ""
		logger.colorError = ""
		logger.colorFatal = ""
	}

	return &logger
}

// SetTimeFormatString sets the format string for the time
func (l *Logger) SetTimeFormatString(formatString string) {
	l.timeFormatString = formatString
}

// LogDebug logs a debug message
func (l *Logger) LogDebug(message string, args ...interface{}) {
	if l.level < 3 {
		return
	}

	funcName := l.getFunctionName()

	outStr, nonColorString := l.formatString(l.colorDebug, "DBG", funcName, fmt.Sprintf(message, args...))

	_, _ = fmt.Print(outStr)

	_ = l.logToFile(nonColorString)
}

// LogInfo logs an info message
func (l *Logger) LogInfo(message string, args ...interface{}) {
	if l.level < 2 {
		return
	}

	funcName := l.getFunctionName()

	outStr, nonColorString := l.formatString(l.colorInfo, "NFO", funcName, fmt.Sprintf(message, args...))

	_, _ = fmt.Print(outStr)

	_ = l.logToFile(nonColorString)
}

// LogError logs an error message
func (l *Logger) LogError(message string, args ...interface{}) {
	if l.level < 1 {
		return
	}

	funcName := l.getFunctionName()

	outStr, nonColorString := l.formatString(l.colorError, "ERR", funcName, fmt.Sprintf(message, args...))

	_, _ = fmt.Print(outStr)

	_ = l.logToFile(nonColorString)
}

// LogFatal logs a fatal message
func (l *Logger) LogFatal(message string, args ...interface{}) {
	funcName := l.getFunctionName()

	outStr, nonColorString := l.formatString(l.colorFatal, "FAT", funcName, fmt.Sprintf(message, args...))

	_, _ = fmt.Print(outStr)

	_ = l.logToFile(nonColorString)

	os.Exit(1)
}
