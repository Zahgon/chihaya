// Package log adds a thin wrapper around logrus to improve non-debug logging
// performance.
package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var (
	l     = logrus.New()
	debug = false
)

// SetDebug controls debug logging.
func SetDebug(to bool) { _ = "STUB: not implemented"; return }

// SetFormatter sets the formatter.
func SetFormatter(to logrus.Formatter) {
	_ = "STUB: not implemented"

	// SetOutput sets the output.
	return
}

func SetOutput(to io.Writer) {
	_ = "STUB: not implemented"

	// Fields is a map of logging fields.
	return
}

type Fields map[string]interface{}

// LogFields implements Fielder for Fields.
func (f Fields) LogFields() Fields {
	_ = "STUB: not implemented"

	// A Fielder provides Fields via the LogFields method.
	return *new(Fields)
}

type Fielder interface {
	LogFields() Fields
}

// err is a wrapper around an error.
type err struct {
	e error
}

// LogFields provides Fields for logging.
func (e err) LogFields() Fields { _ = "STUB: not implemented"; return *new(Fields) }

// Err is a wrapper around errors that implements Fielder.
func Err(e error) Fielder {
	_ = "STUB: not implemented"

	// mergeFielders merges the Fields of multiple Fielders.
	// Fields from the first Fielder will be used unchanged, Fields from subsequent
	// Fielders will be prefixed with "%d.", starting from 1.
	//
	// must be called with len(fielders) > 0
	return *new(Fielder)
}

func mergeFielders(fielders ...Fielder) logrus.Fields {
	_ = "STUB: not implemented"
	return *new(logrus.Fields)
}

// Debug logs at the debug level if debug logging is enabled.
func Debug(v interface{}, fielders ...Fielder) { _ = "STUB: not implemented"; return }

// Info logs at the info level.
func Info(v interface{}, fielders ...Fielder) { _ = "STUB: not implemented"; return }

// Warn logs at the warning level.
func Warn(v interface{}, fielders ...Fielder) { _ = "STUB: not implemented"; return }

// Error logs at the error level.
func Error(v interface{}, fielders ...Fielder) { _ = "STUB: not implemented"; return }

// Fatal logs at the fatal level and exits with a status code != 0.
func Fatal(v interface{}, fielders ...Fielder) { _ = "STUB: not implemented"; return }
