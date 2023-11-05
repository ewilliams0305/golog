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
### Configure
`golog` requires configuration and to faciliate this the builder pattern is used. 


### Logger

### Writer

