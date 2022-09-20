package enum

type LogLevel uint32

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)
