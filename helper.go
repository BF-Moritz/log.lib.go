package loglibgo

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var pointerRegex = regexp.MustCompile(`^\(\*[a-zA-Z0-9_-]+\)$`)

func (l *Logger) formatString(color, level, function, msg string) (colorString, nonColorString string) {

	timeString := ""
	if l.showTime {
		timeString = fmt.Sprintf("[%s]: ", time.Now().Format(l.timeFormatString))
	}

	colorString = fmt.Sprintf("%s%s[%s] %s: %s%s\n", color, timeString, level, function, msg, l.colorReset)
	nonColorString = fmt.Sprintf("%s[%s] %s: %s\n", timeString, level, function, msg)

	return
}

func (l *Logger) logToFile(msg string) error {
	if l.file == nil {
		return nil
	}

	_, err := l.file.WriteString(msg)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) getFunctionName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	// get function name and split by dot
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	if len(parts) == 0 {
		return ""
	}

	// get last 2 parts, only those can be the struct name and the function name
	parts = parts[len(parts)-2:]
	if strings.Contains(parts[0], "/") {
		parts = parts[1:]
	}

	if len(parts) == 0 {
		return ""
	}

	if pointerRegex.MatchString(parts[0]) {
		// remove (* and ) if the function is a pointer
		parts[0] = parts[0][2 : len(parts[0])-1]
	}

	return strings.Join(parts, ".")
}
