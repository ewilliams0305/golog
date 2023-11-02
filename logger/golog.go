package golog

import (
	"time"
)

type LogEvent struct {
	timestamp time.Time
	verbosity verbosity
	message   string
	props     properties
}

type properties map[string]string

type verbosity string

type Logger interface {
	Verbose(message string, props properties)
	//Debug(message string, props properties)
	//Information(message string, props properties)
	//Warn(message string, props properties)
	//Error(message string, err error, props properties)
	//Fatal(message string, err error, props properties)
}

func (gl *GoLog) Verbose(message string, props properties) {

	gl.write(message, props)
	// for s,i := range gl.sinks{

	// }
}

type GoLog struct {
	sinks []SinkWriter
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

type FmtPrinter struct {
	prefix string
}

func (f *FmtPrinter) WriteTo(message LogEvent) {
	print(message.message)
}

type SinkWriter interface {
	WriteTo(message LogEvent)
}
