package golog

type configuration struct {
	level  LogLevel
	format formatter
}

type formatter string
