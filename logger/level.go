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

//type verbosity string
