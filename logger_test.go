package logging_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/ao-concepts/logging"
	"github.com/stretchr/testify/assert"
)

type writer struct {
	Logs []string
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.Logs = append(w.Logs, string(p))
	return len(p), nil
}

func TestLogger_Output(t *testing.T) {
	assert := assert.New(t)

	// debug
	value := "debug"
	testMessage(assert, func(log logging.Logger) {
		log.ErrDebug(fmt.Errorf(value))
	}, value)

	testMessage(assert, func(log logging.Logger) {
		log.Debug(value)
	}, value)

	// info
	value = "info"
	testMessage(assert, func(log logging.Logger) {
		log.ErrInfo(fmt.Errorf(value))
	}, value)

	testMessage(assert, func(log logging.Logger) {
		log.Info(value)
	}, value)

	// warn
	value = "warn"
	testMessage(assert, func(log logging.Logger) {
		log.ErrWarn(fmt.Errorf(value))
	}, value)

	testMessage(assert, func(log logging.Logger) {
		log.Warn(value)
	}, value)

	// error
	value = "error"
	testMessage(assert, func(log logging.Logger) {
		log.ErrError(fmt.Errorf(value))
	}, value)

	testMessage(assert, func(log logging.Logger) {
		log.Error(value)
	}, value)
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
	l = logging.New(logging.Debug, &w)

	l.Write([]byte("write-test"))
	assert.Len(w.Logs, 1)
	var msg logMessage

	assert.Nil(json.Unmarshal([]byte(w.Logs[0]), &msg))
	assert.Equal("info", msg.Level)
	assert.Equal("write-test", msg.Message)
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

type logMessage struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func testMessage(assert *assert.Assertions, fn func(logging.Logger), value string) {
	r, w, err := os.Pipe()
	assert.Nil(err)
	l := logging.New(logging.Debug, w)

	out := make(chan []byte)

	go func() {
		o, err := ioutil.ReadAll(r)
		assert.Nil(err)
		fmt.Println(string(o))
		out <- o
	}()

	fn(l)
	assert.Nil(w.Close())

	var msg logMessage
	assert.Nil(json.Unmarshal(<-out, &msg))
	assert.Equal(value, msg.Level)
	assert.Equal(value, msg.Message)
}
