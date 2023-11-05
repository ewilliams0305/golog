package golog

type configuration struct {
	level  LogLevel
	format formatter
}

type formatter string

func LoggingConfiguration() loggerConfiguration {

	return &goLog{}
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

func (gl *goLog) Configure(minimuLevel LogLevel) loggerWriter {

	// Do the setup of the required internals
	gl.configuration = configuration{
		level:  minimuLevel,
		format: "",
	}
	return gl
}

func (gl *goLog) WriteTo(sink SinkWriter) createWriters {

	config := sinkConfiguration{sink: sink, level: gl.level, template: ""}
	gl.sinks = append(gl.sinks, config)
	return gl
}

func (gl *goLog) MinimuLevel(level LogLevel) createWriters {
	// TODO: Add level restriction to the SINK
	return gl
}

func (gl *goLog) WithFormat(messageTemplate string) createWriters {
	// TODO: Add Message Template to the SINK
	return gl
}

func (gl *goLog) CreateLogger() Logger {

	return gl
}
