package fmtsink

import (
	"fmt"

	golog "github.com/ewilliams0305/golog/logger"
)

type FmtPrinter struct {
}

func (f *FmtPrinter) WriteTo(message golog.LogEvent) error {
	_, e := fmt.Println(colorizeLevel(&message), RenderMessage(&message))
	return e
}

func RenderMessage(e *golog.LogEvent) string {

	if e.Error != nil {
		return e.RenderErrorEvent()
	}

	if len(e.Args) > 0 {
		formattedArgs := formatTemplate(e.Message, e.Args...)
		return fmt.Sprintf(" %v %s", e.Timestamp.Format("2006-01-02T15:04:05"), formattedArgs)
	}
	return fmt.Sprintf(" %v %s", e.Timestamp.Format("2006-01-02T15:04:05"), e.Message)
}

func formatTemplate(template string, args ...any) string {
	return fmt.Sprintf(template, args...)
}

func RenderErrorEvent(e *golog.LogEvent) string {
	if len(e.Args) > 0 {

		formattedArgs := formatTemplate(e.Message, e.Args...)
		return fmt.Sprintf(" %v %s \n%s", e.Timestamp.Format("2006-01-02T15:04:05"), formattedArgs, colorizeError(e.Error))
	}
	return fmt.Sprintf(" %v %s \n%s", e.Timestamp.Format("2006-01-02T15:04:05"), e.Message, colorizeError(e.Error))
}

func colorizeError(err *error) string {
  return ">>>>> " red + error + reset
}

func colorizeLevel(e *golog.LogEvent) string {
	switch e.Level {
	case golog.Verbose, golog.Debug:
		return green + "[" + e.Level.ToString() + "]" + reset
	case golog.Information, golog.Warn:
		return yellow + "[" + e.Level.ToString() + "]" + reset
	case golog.Error, golog.Fatal:
		return red + "[" + e.Level.ToString() + "]" + reset
	default:
		return red + "[" + e.Level.ToString() + "]" + reset
	}
}
