package fmtwriter

import (
	"fmt"

	golog "github.com/ewilliams0305/golog/logger"
)

// The FmtPrinter is a golog sink created for simple verbose printing to the std output.
// Messages will be formatted and published to the console.
//
// Usage Example: Create a loger and provide an instance of the FmtPrinter
//
// logger := golog.LoggingConfiguration().
//
//	Configure(golog.Verbose, "[%l %t] %m").
//	WriteTo(&fmtwriter.FmtPrinter{}).MinimuLevel(golog.Verbose).
//	WriteTo(sink2).MinimuLevel(golog.Information).
//	CreateLogger()
//
// logger.Verbose("Verbose Message %s", "VERNON")
type FmtPrinter struct {
}

// Writes to the console after formatting the log event.
func (f *FmtPrinter) WriteTo(message golog.LogEvent) error {
	_, e := fmt.Println(colorizeLevel(&message), renderMessage(&message))
	return e
}

func renderMessage(e *golog.LogEvent) string {

	if e.Error != nil {
		return renderErrorEvent(e)
	}

	if len(e.Args) > 0 {
		formattedArgs := formatTemplate(e.Message, e.Args...)
		return fmt.Sprintf("%v %s", e.Timestamp.Format("2006-01-02T15:04:05"), formattedArgs)
	}
	return fmt.Sprintf("%v %s", e.Timestamp.Format("2006-01-02T15:04:05"), e.Message)
}

func formatTemplate(template string, args ...any) string {
	return fmt.Sprintf(template, args...)
}

func renderErrorEvent(e *golog.LogEvent) string {
	if len(e.Args) > 0 {

		formattedArgs := formatTemplate(e.Message, e.Args...)
		return fmt.Sprintf("%v %s \n%s", e.Timestamp.Format("2006-01-02T15:04:05"), formattedArgs, colorizeError(e.Error))
	}
	return fmt.Sprintf("%v %s \n%s", e.Timestamp.Format("2006-01-02T15:04:05"), e.Message, colorizeError(e.Error))
}
