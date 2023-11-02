package golog

type SinkWriter interface {
	WriteTo(message LogEvent)
}

type FmtPrinter struct {
	//prefix string
}

func (f FmtPrinter) WriteTo(message LogEvent) {
	println(message.message)
}
