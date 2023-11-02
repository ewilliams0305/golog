package golog

type SinkWriter interface {
	WriteTo(message LogEvent)
}
