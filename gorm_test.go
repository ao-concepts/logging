package logging_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ao-concepts/logging"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

type writer struct {
	Logs []string
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.Logs = append(w.Logs, string(p))
	return len(p), nil
}

func TestCreateGormLogger(t *testing.T) {
	assert := assert.New(t)

	w := writer{}
	l := logging.New(logging.Debug, &w)

	gl := l.CreateGormLogger()

	// logMode
	assert.NotNil(gl.LogMode(logger.Info))

	// info
	gl.Info(context.Background(), "%s", "info")
	assert.Len(w.Logs, 1)
	assert.Equal("{\"level\":\"info\",\"message\":\"info\"}\n", w.Logs[0])

	// warn
	gl.Warn(context.Background(), "%s", "warn")
	assert.Len(w.Logs, 2)
	assert.Equal("{\"level\":\"warn\",\"message\":\"warn\"}\n", w.Logs[1])

	// error
	gl.Error(context.Background(), "%s", "error")
	assert.Len(w.Logs, 3)
	assert.Equal("{\"level\":\"error\",\"message\":\"error\"}\n", w.Logs[2])

	// trace error
	gl.Trace(context.Background(), time.Now(), func() (string, int64) {
		return "", 0
	}, fmt.Errorf("trace"))
	assert.Len(w.Logs, 4)
	assert.Equal("{\"level\":\"error\",\"error\":\"trace\",\"message\":\"trace\"}\n", w.Logs[3])

	// trace 0 rows
	before := time.Now()
	time.Sleep(5 * time.Millisecond)
	gl.Trace(context.Background(), before, func() (string, int64) {
		return "value", 0
	}, nil)
	assert.Len(w.Logs, 5)

	time.Sleep(5 * time.Millisecond)
	gl.Trace(context.Background(), before, func() (string, int64) {
		return "value", -1
	}, nil)
	assert.Len(w.Logs, 6)
}
