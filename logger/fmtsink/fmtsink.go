package fmtsink

import (
	"fmt"
	"strings"

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

func colorizeError(err error) string {
	var sb strings.Builder
	sb.WriteString(">>>>> ")
	sb.WriteString(red)
	sb.WriteString(err.Error())
	sb.WriteString(reset)
	return sb.String()
}

func getLevelColor(l *golog.LogLevel) string {
	switch *l {
	case golog.Verbose, golog.Debug:
		return green
	case golog.Information, golog.Warn:
		return yellow
	case golog.Error, golog.Fatal:
		return red
	default:
		return red
	}
}
func colorizeLevel(e *golog.LogEvent) string {
	color := getLevelColor(&e.Level)
	var sb strings.Builder

	sb.WriteString(color)
	sb.WriteString("[")
	sb.WriteString(e.Level.ToString())
	sb.WriteString("]")
	sb.WriteString(reset)

	return sb.String()
}
