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

	logger.Verbose("Verbose Message %s", "VERNON")
	logger.Debug("Debug Message %s %d", "BILLY", 20)
	logger.Information("Information Message %s", "IMAC")
	logger.Warn("Warn Message %s %d", []interface{}{"Alice", 30})
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

func (f FmtPrinter) WriteTo(message golog.LogEvent) error {

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

func (f FmtPrinterSlow) WriteTo(message golog.LogEvent) error {

	duration := 2 * time.Second
	time.Sleep(duration)
	_, e := fmt.Println(message.RenderMessage())
	return e
}
