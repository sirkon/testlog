package testlog

// TestingPrinter wrapper over *testing.T to print data
type TestingPrinter interface {
	Log(a ...any)
	Logf(format string, a ...any)
	Error(a ...any)
	Errorf(format string, a ...any)
}
