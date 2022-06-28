package testlog_test

import (
	stderrs "errors"
	"github.com/sirkon/errors"
	"github.com/sirkon/testlog"
	"testing"
)

func TestLogging(t *testing.T) {
	t.Run("log-std-error", func(t *testing.T) {
		testlog.Log(t, stderrs.New("not an error"))
	})

	t.Run("log-ctxed-error", func(t *testing.T) {
		testlog.Log(t, errors.New("ctx error").Int("int", 12).Any("map", map[string]string{
			"a": "b",
		}).Str("string", "str"))
	})

	t.Run("error", func(t *testing.T) {
		testlog.Error(t, errors.New("error").Bool("is-error", true))
	})
}
