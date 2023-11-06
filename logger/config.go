package golog

// A struct to store all configuration options used by golog and sink writers.
// The configuration determines the LogLevel and message template of the sink and or globally.
type configuration struct {
	// The golog's configuration logging level.
	// A sink with a specified log level will override the log level configured globally.
	level LogLevel
	// A text formatting template.
	// This is TBD but you will be able to specify some sort of template to dectate how messages are rendered.
	// Something like: [HH:MM:SS LVL] MESSAGE \n ERROR
	format formatter
}

// The type used to define a message template.
type formatter string
