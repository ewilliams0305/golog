package main

import (
	"errors"
	"fmt"

	golog "github.com/ewilliams0305/golog/logger"
)

func main() {

	sink := &FmtPrinter{}

	logger := golog.LoggingConfiguration().
		Configure(golog.Information).
		WriteTo(sink).
		WriteTo(sink).
		CreateLogger()

	logger.Verbose("Verbose Message", nil)
	logger.Debug("Debug Message", nil)
	logger.Information("Information Message", nil)
	logger.Warn("Warn Message", nil)
	logger.Error("Error Message", errors.New("ERROR"), nil)
	logger.Fatal("Fatal Message", errors.New("FATAL"), nil)

}

// type FmtPrinter struct {
// 	//prefix string
// }

// func (f FmtPrinter) WriteTo(message golog.LogEvent) {
// 	//println (message)
// }

type FmtPrinter struct {
	//prefix string
}

func (f FmtPrinter) WriteTo(message golog.LogEvent) {
	fmt.Println(message.RenderMessage())
}
