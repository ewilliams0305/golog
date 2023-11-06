package main

import (
	"errors"
	"fmt"
	"time"

	golog "github.com/ewilliams0305/golog/logger"
	fmtwriter "github.com/ewilliams0305/golog/writers/fmtsink"
)

func main() {

	sink2 := &FmtPrinterSlow{}

	logger := golog.LoggingConfiguration().
		Configure(golog.Verbose, "[%l %t] %m").
		WriteTo(&fmtwriter.FmtPrinter{}).MinimuLevel(golog.Verbose).
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

	newError := errors.New("ERROR MESSAGE PASSED TO LOG")

	logger.Verbose("Verbose Message %s", "VERNON")
	logger.Debug("Debug Message %s %d", "BILLY", 20)
	logger.Information("Information Message %s", "IMAC")
	logger.Warn("Warn Message %s %d", "Alice", 30)
	logger.Error("Error Message", newError)
	logger.Fatal("Fatal Message", newError)

}

func formatTemplate(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
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
