/*
The golog package is a logging framework designed around abstracting the output if logging messages. 
The golog framework helps to: 
- Describe a log message with structure 
- Render the messgae
- Abstract the output of the messages to different and or multiple destinations. 

Example: BYOL bring your own logger

func main() {

	sink1 := &FmtPrinter{}
	sink2 := &FmtPrinter{}

	logger := golog.LoggingConfiguration().
		Configure(golog.Information).
		WriteTo(sink1).MinimuLevel(golog.Debug).WithFormat("").
		WriteTo(sink2).MinimuLevel(golog.Information).
		CreateLogger()

	logger.Verbose("Verbose Message", nil)
	logger.Debug("Debug Message", nil)
	logger.Information("Information Message", nil)
	logger.Warn("Warn Message", nil)
	logger.Error("Error Message", errors.New("ERROR"), nil)
	logger.Fatal("Fatal Message", errors.New("FATAL"), nil)

}

*/
package golog

import (
	"time"
)

// The GoLog struc is the core component of the golog package. 
// GoLog is responsible for storing your logging confogurations, 
// pointers to your sinks, and implementations of the loger builder.
// While tou will never directly access the golog it a critical component of the framework. 
type GoLog struct {
	sinks []SinkWriter
	configuration
}

// The logger is NOT responsible for writing messages to sinks. 
// In fact a logger simply passes data to the golog and LogEvents are created. 
//
// When a Logger method is invoked these events are passed to the sinks and rendered for display. 
// The logger is Generate by the builder pattern and then used throughout your application.  
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

	c := make(chan string)
	for i, s := range gl.sinks {
		//wg.Add(1)

		go writeSink(s, LogEvent{
			timestamp: time.Now(),
			level:     level,
			message:   message,
			props:     props,
		}, c)

		// go s.WriteTo(LogEvent{
		// 	timestamp: time.Now(),
		// 	level:     level,
		// 	message:   message,
		// 	props:     props,
		// })

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

func writeSink(s SinkWriter, e LogEvent, c chan string) {
	s.WriteTo(e)

}
