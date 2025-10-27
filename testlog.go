package testlog

import (
	"fmt"
	"go/token"
	"runtime"
	"strings"

	"github.com/sirkon/errors"
)

const (
	bold = "\033[1m"
	red  = "\033[1;31m"
)

// TestLog logger with stored TestingPrinter.
type TestLog struct {
	t TestingPrinter
}

// New creates TestLogInstance.
func New(t TestingPrinter) TestLog {
	return TestLog{
		t: t,
	}
}

// Log logs an error.
func (tl TestLog) Log(err error) {
	_, _ = fmt.Fprintln(tl.t.Output(), renderString(err, bold))
}

// Error signal an error.
func (tl TestLog) Error(err error) {
	_, _ = fmt.Fprintln(tl.t.Output(), renderString(err, red))
}

// Log logs error.
func Log(t TestingPrinter, err error) {
	_, _ = fmt.Fprintln(t.Output(), renderString(err, bold))
}

// Error signal error.
func Error(t TestingPrinter, err error) {
	_, _ = fmt.Fprintln(t.Output(), renderString(err, red))
}

// Check do nothing and return false if error is nil.
// Prints error and return true otherwise.
func Check(t TestingPrinter, err error) bool {
	if err == nil {
		return false
	}

	_, _ = fmt.Fprintln(t.Output(), renderString(err, red))
	return true
}

func renderString(err error, highlight string) string {
	_, fn, line, _ := runtime.Caller(2)
	pos := token.Position{Filename: fn, Line: line}

	if err == nil {
		return pos.String() + ": <nil>"
	}

	var b strings.Builder
	b.WriteString(pos.String())
	b.WriteByte(' ')
	b.WriteString(highlight)
	b.WriteString(err.Error())
	b.WriteString("\033[0m")

	d := errors.GetContextDeliverer(err)
	if d == nil {
		return b.String()
	}

	var c errorContextConsumer
	d.Deliver(&c)

	b.WriteByte('\n')
	origIdent := "    "
	for _, level := range c.levels {
		ident := origIdent
		b.WriteString(ident)
		b.WriteString(level.what)
		b.WriteByte('\n')

		ident += "  "
		if level.loc != "" {
			b.WriteString(ident)
			b.WriteString("@location: ")
			b.WriteString(level.loc)
			b.WriteByte('\n')
		}
		for _, v := range level.vars {
			b.WriteString(ident)
			b.WriteString(v.name)
			b.WriteString(": ")
			b.WriteString(fmt.Sprint(v.value))
			b.WriteByte('\n')
		}
	}

	return b.String()
}
