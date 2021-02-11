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

func TestCreateGormLogger(t *testing.T) {
	assert := assert.New(t)

	// logMode
	l := logging.New(logging.Debug, nil)
	gl := l.CreateGormLogger()
	assert.NotNil(gl.LogMode(logger.Info))

	// info
	testMessage(assert, func(log logging.Logger) {
		gl := log.CreateGormLogger()
		gl.Info(context.Background(), "%s", "info")
	}, "info")

	// warn
	testMessage(assert, func(log logging.Logger) {
		gl := log.CreateGormLogger()
		gl.Warn(context.Background(), "%s", "warn")
	}, "warn")

	// // error
	testMessage(assert, func(log logging.Logger) {
		gl := log.CreateGormLogger()
		gl.Error(context.Background(), "%s", "error")
	}, "error")

	// trace error
	testMessage(assert, func(log logging.Logger) {
		gl := log.CreateGormLogger()
		gl.Trace(context.Background(), time.Now(), func() (string, int64) {
			return "", 0
		}, fmt.Errorf("error"))
	}, "error")

	// trace 0 rows
	w := &writer{}
	log := logging.New(logging.Debug, w)
	glog := log.CreateGormLogger()

	before := time.Now()
	time.Sleep(5 * time.Millisecond)
	glog.Trace(context.Background(), before, func() (string, int64) {
		return "value", 0
	}, nil)
	assert.Len(w.Logs, 1)

	time.Sleep(5 * time.Millisecond)
	glog.Trace(context.Background(), before, func() (string, int64) {
		return "value", -1
	}, nil)
	assert.Len(w.Logs, 2)
}
