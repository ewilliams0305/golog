package main

import (
	"errors"
	"fmt"
	"time"

	golog "github.com/ewilliams0305/golog/logger"
)

func main() {

	sink1 := &FmtPrinter{}
	sink2 := &FmtPrinterSlow{}

	logger := golog.LoggingConfiguration().
		Configure(golog.Verbose, "[%l %t] %m").
		WriteTo(sink1).MinimuLevel(golog.Verbose).
		WriteTo(sink2).MinimuLevel(golog.Information).
		CreateLogger()

	// GOLOG provide you runtime log level switching.
	// When log levels are switched at runtime it will syncronize and revert the level of all sinks.
	// In the future you will be able to switch the level of each sink independantly.
	var response string

	fmt.Print("Enter a new Log Level (verbose, debug, info, warn, error, fatal): ")
	_, err := fmt.Scanln(&response)

	if err != nil {
		logger.Error("Proceeding with default levels", err)
	} else {

	}
	level := golog.CreateLevelFromString(response)
	logger.SwitchLevel(level)

	logger.Verbose("Verbose Message %s", "VERNON")
	logger.Debug("Debug Message %s %d", "BILLY", 20)
	logger.Information("Information Message %s", "IMAC")
	logger.Warn("Warn Message %s %d", "Alice", 30)
	logger.Error("Error Message", errors.New("ERROR"))
	logger.Fatal("Fatal Message", errors.New("FATAL"))

}

func formatTemplate(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}

/***************************
*
* Mock Logger that is FAST
*
****************************/

type FmtPrinter struct {
}

func (f *FmtPrinter) WriteTo(message golog.LogEvent) error {

	_, e := fmt.Println(message.RenderMessage())
	return e
}

/***************************
*
* Mock Logger that is FAST
*
****************************/

type FmtPrinterSlow struct {
}

func (f *FmtPrinterSlow) WriteTo(message golog.LogEvent) error {

	duration := 2 * time.Second
	time.Sleep(duration)
	_, e := fmt.Println(message.RenderMessage())
	return e
}
