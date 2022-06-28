package testlog

import (
	"encoding/json"
	"github.com/sirkon/errors"
	"strings"
)

const (
	bold = "\033[1m"
	red  = "\033[1;31m"
)

// Log logs error.
func Log(t TestingPrinter, err error) {
	t.Log(renderString(err, bold))
}

// Error signal error.
func Error(t TestingPrinter, err error) {
	t.Error(renderString(err, red))
}

// Check do nothing and return false if error is nil.
// Prints error and return true otherwise.
func Check(t TestingPrinter, err error) bool {
	if err == nil {
		return false
	}

	t.Error(renderString(err, red))
	return true
}

func renderString(err error, highlight string) string {
	if err == nil {
		return "<nil>"
	}

	var b strings.Builder
	b.WriteString(highlight)
	b.WriteString(err.Error())
	b.WriteString("\033[0m")

	d := errors.GetContextDeliverer(err)
	if d == nil {
		return b.String()
	}

	var c errorContextConsumer
	d.Deliver(&c)

	if len(c.vars) == 0 {
		return b.String()
	}

	b.WriteString("\t{")
	for i, v := range c.vars {
		if i > 0 {
			b.WriteString(", ")
		}

		b.WriteByte('"')
		b.WriteString(v.name)
		b.WriteString(`":`)
		val, err := json.Marshal(v.value)
		if err != nil {
			panic(errors.Wrap(err, "marshal context value "+v.name))
		}
		b.Write(val)
	}
	b.WriteByte('}')

	return b.String()
}
