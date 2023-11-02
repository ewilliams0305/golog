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
}

type loggerWriter interface {
	WriteTo(sink SinkWriter) createWriters
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
	// Add the sink to the writers

	gl.sinks = append(gl.sinks, sink)
	return gl
}

func (gl *GoLog) CreateLogger() Logger {

	return gl
}
