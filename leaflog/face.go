package leaflog

type ILogger interface {
	Debug(format string, a ...interface{})
	Release(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	// Close()
}
