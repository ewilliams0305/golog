package golog

// The entry point for the [golog] logging framework.
// The [LoggingConfiguration] functions starts the golog configuration builder by returning a [loggerConfiguration]
// interface.  The logging configuration proceeds the consumer to the next step configuring an log level and format template.
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

// Overrides the globally configured log level with a sink specific minimum level.
// This allow several sinks to have a log level of verbose while file system or http sinks onyl log errors.
type writerLevel interface {
	MinimuLevel(level LogLevel) createWriters
}

// Overrides the globally configured message format template.
// When overriden specific sinks can use a custom format specific to the sink.
type writerFormat interface {
	WithFormat(format formatter) createWriters
}

// The final step in the builder process returns a reference to the golog instance as a Logger interface.
// The logger interface is now configured and contains all required function to log data to the logging sinks.
type createLogger interface {
	CreateLogger() Logger
}

/******************************************************************************************
* Builder interface implemenations.
* All functions below pass the golog struct between each call augmenting it with additional
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

	writer := &loggingSink{
		sink: &sink,
		config: configuration{
			level:  gl.config.level,
			format: gl.config.format,
		},
	}
	gl.sinks = append(gl.sinks, writer)

	gl.sinkIndex++
	return gl
}

// Overrides the globally configured log level with a sink specific minimum level.
// This allow several sinks to have a log level of verbose while file system or http sinks onyl log errors.
func (gl *goLog) MinimuLevel(level LogLevel) createWriters {

	gl.sinks[gl.sinkIndex-1].config.level = level
	return gl
}

// Overrides the globally configured message format template.
// When overriden specific sinks can use a custom format specific to the sink.
func (gl *goLog) WithFormat(format formatter) createWriters {
	gl.sinks[gl.sinkIndex-1].config.format = format
	return gl
}

func (gl *goLog) CreateLogger() Logger {

	return gl
}
