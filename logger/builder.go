package golog

func LoggingConfiguration() loggerConfiguration {

	return &goLog{}
}

type loggerConfiguration interface {
	Configure(minimuLevel LogLevel, template string) loggerWriter
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
	WithFormat(format formatter) createWriters
}

type createLogger interface {
	CreateLogger() Logger
}

/******************************************************************************************
* Builder interface implemenations.
* All functions below pass the golog struct between each call augmenting it with addtional
* behavior.  Once the builder is completed the golog struct is used to map messages to sinks.
*******************************************************************************************/

func (gl *goLog) Configure(minimuLevel LogLevel, template string) loggerWriter {

	// Do the setup of the required internals
	gl.config = configuration{
		level:  minimuLevel,
		format: "",
	}
	return gl
}

func (gl *goLog) WriteTo(sink SinkWriter) createWriters {

	config := loggingSink{
		sink: sink,
		config: configuration{
			level:  gl.config.level,
			format: gl.config.format,
		},
	}
	gl.sinks = append(gl.sinks, config)
	return gl
}

func (gl *goLog) MinimuLevel(level LogLevel) createWriters {
	// TODO: Add level restriction to the SINK
	return gl
}

func (gl *goLog) WithFormat(format formatter) createWriters {
	// TODO: Add Message Template to the SINK
	return gl
}

func (gl *goLog) CreateLogger() Logger {

	return gl
}
