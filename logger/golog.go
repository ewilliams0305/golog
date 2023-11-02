package golog

import (
	"time"
)

type properties map[string]string

type Logger interface {
	Verbose(message string, props properties)
	Debug(message string, props properties)
	Information(message string, props properties)
	Warn(message string, props properties)
	Error(message string, err error, props properties)
	Fatal(message string, err error, props properties)
}

func (gl *GoLog) Verbose(message string, props properties) {

	if gl.configuration.level <= Verbose {
		gl.write(message, Verbose, props)
	}
}

func (gl *GoLog) Debug(message string, props properties) {

	if gl.configuration.level <= Debug {
		gl.write(message, Debug, props)
	}
}

func (gl *GoLog) Information(message string, props properties) {

	if gl.configuration.level <= Information {
		gl.write(message, Information, props)
	}
}

func (gl *GoLog) Warn(message string, props properties) {

	if gl.configuration.level <= Warn {
		gl.write(message, Warn, props)
	}
}

func (gl *GoLog) Error(message string, err error, props properties) {

	if gl.configuration.level <= Error {
		gl.write(message, Error, props)
	}
}

func (gl *GoLog) Fatal(message string, err error, props properties) {

	if gl.configuration.level <= Fatal {
		gl.write(message, Fatal, props)
	}
}

type GoLog struct {
	sinks []SinkWriter
	configuration
}

func (gl *GoLog) write(message string, level LogLevel, props properties) {

	//var wg sync.WaitGroup

	for i, s := range gl.sinks {
		//wg.Add(1)

		s.WriteTo(LogEvent{
			timestamp: time.Now(),
			level:     level,
			message:   message,
			props:     props,
		})

		// go func(index int, sink SinkWriter) {

		// 	//defer wg.Done()

		// 	sink.WriteTo(LogEvent{
		// 		timestamp: time.Now(),
		// 		level:     level,
		// 		message:   message,
		// 		props:     props,
		// 	})
		// }(i, s)

		println(s, i)
	}
}
