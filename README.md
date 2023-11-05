[![Go Report Card](https://goreportcard.com/badge/github.com/ewilliams0305/golog)](https://goreportcard.com/report/github.com/ewilliams0305/golog)

# golog
Logging abstraction for go. 

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

### Logger

### Writer

