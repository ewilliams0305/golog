package golog

import "strings"

// The LogLevel type represents a restriction put on a writers output
// The LogLevel supports a total of 6 levels.
// When a log level is set to Information (2) only log message of
// Information or higher will be generated.
type LogLevel uint8

// The default level values used by the LogLevel
// type. All other values will be be defined as unknown and
// log messages will not be generated.
const (
	Verbose     LogLevel = 0
	Debug       LogLevel = 1
	Information LogLevel = 2
	Warn        LogLevel = 3
	Error       LogLevel = 4
	Fatal       LogLevel = 5
)

// Converts the LogLevel type to a string value used in log messages.
// All unknown  alues will be UNKNOWN.
func (s LogLevel) ToString() string {
	switch s {
	case Verbose:
		return "VRB"
	case Debug:
		return "DBG"
	case Information:
		return "INF"
	case Warn:
		return "WRN"
	case Error:
		return "ERR"
	case Fatal:
		return "FTL"
	}
	return "UNKNOWN"
}

// Creates a level from a string value.
func CreateLevelFromString(s string) LogLevel {

	switch strings.ToUpper(s) {
	case "VERBOSE":
		return Verbose
	case "DEBUG":
		return Debug
	case "INFO":
		return Information
	case "WARN":
		return Warn
	case "ERROR":
		return Error
	case "FATAL":
		return Fatal
	}
	return Information
}
