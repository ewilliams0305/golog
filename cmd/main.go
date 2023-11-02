package main

import golog "github.com/ewilliams0305/golog/logger"

func main() {

	sink := golog.FmtPrinter{}

	sink.WriteTo(golog.LogEvent{})

	logger := golog.LoggingConfiguration().
		Create().
		WriteTo(sink)

	logger.Verbose("Verbose Message", nil)

}
