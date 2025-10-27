package testlog

import (
	"io"
)

// TestingPrinter wrapper over *testing.T to print data
type TestingPrinter interface {
	Output() io.Writer
}
