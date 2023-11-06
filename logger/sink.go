package golog

// The SinkWriter is the 'exit point' for log event messages.
// The GoLong struct will store a pointer to all SinkWriters registered with golog.
// Sinks can be registered with golog during the builder process.
// When a log message is sent to a sink the message will only be written when the log level is set lower than the message event.
type SinkWriter interface {

	// The write method will be invoked when a log messge is created.
	// The write method will only be invoked if
	// - The log message is greater than the sinks current level.
	WriteTo(message LogEvent) error
}
