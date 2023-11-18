package enum

type LogLevel uint32

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

var stringMap map[LogLevel]string = map[LogLevel]string{
	LogLevelNone:  "None",
	LogLevelError: "Error",
	LogLevelInfo:  "Info",
	LogLevelDebug: "Debug",
}

func (l LogLevel) ToString() string {
	return stringMap[l]
}

var reverseStringMap map[string]LogLevel = map[string]LogLevel{
	"None":  LogLevelNone,
	"Error": LogLevelError,
	"Info":  LogLevelInfo,
	"Debug": LogLevelDebug,
}

func LogLevelFromString(str string) LogLevel {
	return reverseStringMap[str]
}
