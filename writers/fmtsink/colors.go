package fmtwriter

import (
	"strings"

	golog "github.com/ewilliams0305/golog/logger"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

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
