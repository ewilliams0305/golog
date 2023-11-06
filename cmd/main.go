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
		WriteTo(sink2).MinimuLevel(golog.Verbose).
		CreateLogger()

	logger.Verbose("Verbose Message", nil)
	logger.Debug("Debug Message", nil)
	logger.Information("Information Message", nil)
	logger.Warn("Warn Message", nil)
	logger.Error("Error Message", errors.New("ERROR"), nil)
	logger.Fatal("Fatal Message", errors.New("FATAL"), nil)

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
