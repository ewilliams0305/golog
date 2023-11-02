package main

import golog "github.com/ewilliams0305/golog/logger"

func main() {

	sink := &golog.FmtPrinter{}

	sink.WriteTo(golog.LogEvent{})

	logger := golog.LoggingConfiguration().
		Configure(golog.Verbose).
		WriteTo(sink).
		WriteTo(sink).
		CreateLogger()

	logger.Verbose("Verbose Message", nil)
	logger.Debug("Debug Message", nil)
	logger.Information("Information Message", nil)

}
