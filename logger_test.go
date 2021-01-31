package logging_test

import (
	"fmt"
	"testing"

	"github.com/ao-concepts/logging"
	"github.com/stretchr/testify/assert"
)

func ExampleLogger() {
	l := logging.New(logging.Debug, nil)

	l.ErrDebug(fmt.Errorf("debug"))
	l.Debug("debug")
	l.ErrInfo(fmt.Errorf("info"))
	l.Info("info")
	l.ErrWarn(fmt.Errorf("warn"))
	l.Warn("warn")
	l.ErrError(fmt.Errorf("error"))
	l.Error("error")
	// Output: {"level":"debug","error":"debug","message":"debug"}
	// {"level":"debug","message":"debug"}
	// {"level":"info","error":"info","message":"info"}
	// {"level":"info","message":"info"}
	// {"level":"warn","error":"warn","message":"warn"}
	// {"level":"warn","message":"warn"}
	// {"level":"error","error":"error","message":"error"}
	// {"level":"error","message":"error"}
}

func TestLogger(t *testing.T) {
	assert := assert.New(t)

	// debug level
	w := writer{}
	l := logging.New(logging.Debug, &w)

	runLogs(l, assert)
	assert.Len(w.Logs, 10)

	// info level
	w = writer{}
	l = logging.New(logging.Info, &w)

	runLogs(l, assert)
	assert.Len(w.Logs, 8)

	// warn level
	w = writer{}
	l = logging.New(logging.Warn, &w)

	runLogs(l, assert)
	assert.Len(w.Logs, 6)

	// error level
	w = writer{}
	l = logging.New(logging.Error, &w)

	runLogs(l, assert)
	assert.Len(w.Logs, 4)

	// fatal level
	w = writer{}
	l = logging.New(logging.Fatal, &w)

	runLogs(l, assert)
	assert.Len(w.Logs, 2)

	// write
	w = writer{}
	l = logging.New(logging.Fatal, &w)

	l.Write([]byte("write-test"))
	assert.Len(w.Logs, 1)
	assert.Equal("{\"message\":\"write-test\"}\n", w.Logs[0])
}

func runLogs(l logging.Logger, assert *assert.Assertions) {
	assert.Panics(func() {
		l.Fatal("fatal")
	})
	assert.Panics(func() {
		l.ErrFatal(fmt.Errorf("fatal"))
	})
	l.Error("error")
	l.ErrError(fmt.Errorf("error"))
	l.Warn("warn")
	l.ErrWarn(fmt.Errorf("warn"))
	l.Info("info")
	l.ErrInfo(fmt.Errorf("info"))
	l.Debug("debug")
	l.ErrDebug(fmt.Errorf("debug"))
}
