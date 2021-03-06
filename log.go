package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/BF-Moritz/log.lib.go/src/consts"
)

type Logger struct {
	level            uint32
	showTime         bool
	timeFormatString string

	colorReset string
	colorDebug string
	colorInfo  string
	colorError string
	colorFatal string
}

// NewLogger Creates a new Logger
func NewLogger(level uint32, showTime, showColor bool) *Logger {
	var logger Logger = Logger{
		level:            level,
		showTime:         showTime,
		timeFormatString: consts.TimeFormatString,

		colorReset: consts.Reset,
		colorDebug: consts.Cyan,
		colorInfo:  consts.Green,
		colorError: consts.Red,
		colorFatal: consts.Purple,
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

func (l *Logger) LogDebug(function, message string, args ...interface{}) {
	if l.level < 3 {
		return
	}

	timeString := ""
	if l.showTime {
		timeString = fmt.Sprintf("[%s]: ", time.Now().Format(l.timeFormatString))
	}
	fmt.Printf("%s%s[DBG] %s: %s%s\n", l.colorDebug, timeString, function, fmt.Sprintf(message, args...), l.colorReset)
}

func (l *Logger) LogInfo(function, message string, args ...interface{}) {
	if l.level < 2 {
		return
	}

	timeString := ""
	if l.showTime {
		timeString = fmt.Sprintf("[%s]: ", time.Now().Format(l.timeFormatString))
	}
	fmt.Printf("%s%s[NFO] %s: %s%s\n", l.colorInfo, timeString, function, fmt.Sprintf(message, args...), l.colorReset)
}

func (l *Logger) LogError(function, message string, args ...interface{}) {
	if l.level < 1 {
		return
	}
	timeString := ""
	if l.showTime {
		timeString = fmt.Sprintf("[%s]: ", time.Now().Format(l.timeFormatString))
	}
	fmt.Printf("%s%s[ERR] %s: %s%s\n", l.colorError, timeString, function, fmt.Sprintf(message, args...), l.colorReset)
}

func (l *Logger) LogFatal(function, message string, args ...interface{}) {
	timeString := ""
	if l.showTime {
		timeString = fmt.Sprintf("[%s]: ", time.Now().Format(l.timeFormatString))
	}
	log.Fatalf("%s%s[FAT] %s: %s%s\n", l.colorFatal, timeString, function, fmt.Sprintf(message, args...), l.colorReset)
}
