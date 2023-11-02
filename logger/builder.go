package golog

type loggerConfiguration interface {
	Create() loggerWriter
}

type loggerWriter interface {
	WriteTo(sink *SinkWriter) Logger
}

func LoggingConfiguration() loggerConfiguration {

	return &GoLog{}
}

func (gl *GoLog) Create() loggerWriter {

	// Do the setup of the required internals

	return gl
}

func (gl *GoLog) WriteTo(sink *SinkWriter) Logger {
	// Add the sink to the writers

	s := append(gl.sinks, *sink)
	println(s)

	return gl
}
