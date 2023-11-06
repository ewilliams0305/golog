[![Go Report Card](https://goreportcard.com/badge/github.com/ewilliams0305/golog)](https://goreportcard.com/report/github.com/ewilliams0305/golog)

# golog
A Logging abstraction for go that seperates outputs ```SinkWriter```
from messages ```Logger```. golog can support multiple logging configurations 
and or multiple output configurations. golog supports the ability
to enable and disable verbosity at runtime. golog maintains structure to your logged 
messages to ensure data integrity and ease of queries. 

```go
writer := &FmtPrinter{}

logger := golog.LoggingConfiguration().
	Configure(golog.Information).
	WriteTo(writer).MinimuLevel(golog.Debug).WithFormat("").
	CreateLogger()

logger.Verbose("Verbose Message", nil)
logger.Debug("Debug Message", nil)
logger.Information("Information Message", nil)
logger.Warn("Warn Message", nil)
logger.Error("Error Message", errors.New("ERROR"), nil)
logger.Fatal("Fatal Message", errors.New("FATAL"), nil)

```
## Configure
`golog` requires configuration, to faciliate this the builder pattern is used. 
The golog configuration API was designed to ensure consumers fall into the pit of success. 

The entry point to the golog confoguration builder is the LoggingConfiguration function. 
This function creates a new GoLog struct and returns an implementation of the Builder interface. 
```go
logger := golog.LoggingConfiguration()
```
In doing so consumers are forced to proceed to the bext step of configuration,
Specifying a verbosity level. The configure function accepts one of 6 predefined default
`LogLevel`s; Verbose, Debug, Information, Warn, Error, Fatal. 
Only logging messages with be geberated when the messages LogLevel is >= to the configured LogLevel. 
```go
Configure(golog.Information)
```
Lastly consumers are dirceted to configure the `SinkWriter`s. 
These are destinations to display messages. You can BYOL and create customer sinks (see below) 
or use the golog predefined fmt writer. The WriteTo function allows you to 
provide a Sink, override the LogLevel, and create a message template. The only requirement 
is that at least 1 SinkWriter is provided, templates and overrides are optional. 

```go
WriteTo(writer).MinimuLevel(golog.Debug).WithFormat("")
```
Finally we call CreateLogger and return the logger. 
we can now use out logger to generate messages. When the logger is used messages will be generated for all SinkWriters
and rendered to all destinations with a currently enabled LogLevel. 
```go
CreateLogger()
```
So you have choices; create a single logger for your entire application, or create multiple loggers. a single logger 
can send data to multiple sinks or a single sink. 

## Logger

```go
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
```

## Sink Writers
Creating custom sink writers compatible with golog is as simple as 
implementing the `SinkWriter` interface. 

```go
type SinkWriter interface {
	WriteTo(message LogEvent) error
}
```

Example: the provided /fmtsink implements `SinkWriter`
and passes the data off to a formatting function. 
```go
type FmtPrinter struct {
}

func (f *FmtPrinter) WriteTo(message golog.LogEvent) error {
	_, e := fmt.Println(colorizeLevel(&message), RenderMessage(&message))
	return e
}
```

This new type can now be used during configuration and passed
into the `WriteTo` step. 

Currently the `WriteTo` functions are executed on a go
function. The intent here is to prevent blocking when logging to remote servers. 
In the future this behavior will be configurable during setup. 

