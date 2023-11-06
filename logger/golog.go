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

	logger := golog.Loggingconfig().
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
	"sync"
	"time"
)

// The goLog struc is the core component of the golog package.
// goLog is responsible for storing your logging confogurations,
// pointers to your sinks, and implementations of the loger builder.
// While tou will never directly access the golog it a critical component of the framework.
type goLog struct {
	sinks     []*loggingSink
	config    configuration
	sinkIndex int16
}

type loggingSink struct {
	sink   *SinkWriter
	config configuration
}

// The logger is NOT responsible for writing messages to sinks.
// In fact a logger simply passes data to the golog and LogEvents are created.
//
// When a Logger method is invoked these events are passed to the sinks and rendered for display.
// The logger is Generate by the builder pattern and then used throughout your application.
type Logger interface {
	MessageWriter
	LogSwitch
}

type MessageWriter interface {
	Verbose(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Information(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, err error, args ...interface{})
	Fatal(message string, err error, args ...interface{})
}

type LogSwitch interface {
	SwitchLevel(level LogLevel)
	CurrentLevel() LogLevel
}

func (gl *goLog) SwitchLevel(level LogLevel) {

	gl.config.level = level

	for _, s := range gl.sinks {
		s.config.level = level
	}
}

func (gl *goLog) CurrentLevel() LogLevel {

	return gl.config.level
}

func (gl *goLog) Verbose(message string, args ...interface{}) {

	gl.write(message, Verbose, nil, args...)
}

func (gl *goLog) Debug(message string, args ...interface{}) {

	gl.write(message, Debug, nil, args...)
}

func (gl *goLog) Information(message string, args ...interface{}) {

	gl.write(message, Information, nil, args...)
}

func (gl *goLog) Warn(message string, args ...interface{}) {

	gl.write(message, Warn, nil, args...)
}

func (gl *goLog) Error(message string, err error, args ...interface{}) {

	gl.write(message, Error, err, args...)
}

func (gl *goLog) Fatal(message string, err error, args ...interface{}) {

	gl.write(message, Fatal, err, args...)
}

func (gl *goLog) write(message string, level LogLevel, err error, args ...interface{}) {

	resultChan := make(chan string, len(gl.sinks))
	var wg sync.WaitGroup

	for _, s := range gl.sinks {
		if s.config.level <= level {

			wg.Add(1)

			go func(sink *SinkWriter) {
				writeSink(*sink, LogEvent{
					Timestamp: time.Now(),
					Level:     level,
					Message:   message,
					Error:     err,
					Args:      args,
				}, resultChan)
				defer wg.Done()
			}(s.sink)
		}
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for range resultChan {
	}
}

func writeSink(s SinkWriter, e LogEvent, c chan string) {

	err := s.WriteTo(e)
	if err != nil {
		c <- "Error: " + err.Error()
	} else {
		c <- "Success"
	}
}
