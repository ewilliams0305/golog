package golog

type configuration struct {
	level  LogLevel
	format formatter
}

type formatter string

func LoggingConfiguration() loggerConfiguration {

	return &GoLog{}
}

type loggerConfiguration interface {
	Configure(minimuLevel LogLevel) loggerWriter
}

type createWriters interface {
	loggerWriter
	createLogger
	writerLevel
	writerFormat
}

type loggerWriter interface {
	WriteTo(sink SinkWriter) createWriters
}

type writerLevel interface {
	MinimuLevel(level LogLevel) createWriters
}

type writerFormat interface {
	WithFormat(messageTemplate string) createWriters
}

type createLogger interface {
	CreateLogger() Logger
}

// Function Implementations

func (gl *GoLog) Configure(minimuLevel LogLevel) loggerWriter {

	// Do the setup of the required internals
	gl.configuration = configuration{
		level:  minimuLevel,
		format: "",
	}
	return gl
}

func (gl *GoLog) WriteTo(sink SinkWriter) createWriters {
	gl.sinks = append(gl.sinks, sink)
	return gl
}

func (gl *GoLog) MinimuLevel(level LogLevel) createWriters {
	// TODO: Add level restriction to the SINK
	return gl
}

func (gl *GoLog) WithFormat(messageTemplate string) createWriters {
	// TODO: Add Message Template to the SINK
	return gl
}

func (gl *GoLog) CreateLogger() Logger {

	return gl
}
