package golog

import (
	"time"
)

type properties map[string]string

type Logger interface {
	Verbose(message string, props properties)
	Debug(message string, props properties)
	Information(message string, props properties)
	//Warn(message string, props properties)
	//Error(message string, err error, props properties)
	//Fatal(message string, err error, props properties)
}

func (gl *GoLog) Verbose(message string, props properties) {

	if gl.configuration.level >= Verbose {
		gl.write(message, props)
	}
}

func (gl *GoLog) Debug(message string, props properties) {

	if gl.configuration.level >= Debug {
		gl.write(message, props)
	}
}

func (gl *GoLog) Information(message string, props properties) {

	if gl.configuration.level >= Information {
		gl.write(message, props)
	}
}

type GoLog struct {
	sinks []SinkWriter
	configuration
}

func (gl *GoLog) write(message string, props properties) {
	for i, s := range gl.sinks {
		// MAKE GO FUNC HERE
		s.WriteTo(LogEvent{
			timestamp: time.Now(),
			verbosity: "ERR",
			message:   message,
			props:     props,
		})
		println(s, i)
	}
}
