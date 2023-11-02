package golog

type LogLevel int64

const (
	Verbose     LogLevel = 0
	Debug       LogLevel = 1
	Information LogLevel = 2
	Warn        LogLevel = 3
	Error       LogLevel = 4
	Fatal       LogLevel = 5
)

type verbosity string
