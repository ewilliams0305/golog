package golog

import "time"

type LogEvent struct {
	timestamp time.Time
	verbosity verbosity
	message   string
	props     properties
}
