package golog

import (
	"fmt"
	"time"
)

type LogEvent struct {
	timestamp time.Time
	level     LogLevel
	message   string
	props     properties
}

type FormatMessage interface {
	Format() string
}

func (e *LogEvent) RenderMessage() string {
	return fmt.Sprintf("[%s %v] %s", e.level.ToString(), e.timestamp.Format("2006-01-02T15:04:05"), e.message)
}
