package log

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	stdlog "log"
	"os"
)

var commonOptions, _ = Options(
	WithUTCTimestamp,
	WithMicrosecondsTimestamp,
	WithSourceLocationShort,
	WithLevel(Debug),
)

type stdLogOption interface {
	applyStdLog(*stdLogger) error
}

type stdLogger struct {
	level  Level
	flags  flags
	writer io.WriteCloser
	logger *stdlog.Logger
}

// Close disables and closed the logger, freeing up resources.
func (l *stdLogger) Close() {
	l.level = Disabled
	// stop using the writer before closing it
	l.logger.SetOutput(ioutil.Discard)
	_ = l.writer.Close()
}

func (l *stdLogger) Level() Level {
	return l.level
}

func (l *stdLogger) Logf(level Level, format string, value ...interface{}) {
	if level.IsEnabled(l.level) {
		_ = l.logger.Output(3, fmt.Sprintf(level.String()+": "+format, value...))
	}
}

// New creates a new logger with the specified options.
func New(opt ...Option) (log Log, err error) {
	l := &stdLogger{
		flags: stdlog.LstdFlags,
	}

	// apply default options first
	if err = commonOptions.applyStdLog(l); err != nil {
		err = newConfigError(err)
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.applyStdLog(l); err != nil {
			err = newConfigError(err)
			return
		}
	}

	if l.writer == nil {
		err = newConfigError(errors.New("no writer given"))
		return
	}

	l.logger = stdlog.New(l.writer, "", int(l.flags))

	return Log{logger: l}, nil
}

// NewStderr creates a new logger that logs to stderr. Additional options can be specified using opt.
func NewStderr(opt ...Option) (Log, error) {
	return New(options(opt), WithWriter(os.Stderr))
}

// NewStdout creates a new logger that logs to stdout. Additional options can be specified using opt.
func NewStdout(opt ...Option) (Log, error) {
	return New(options(opt), WithWriter(os.Stdout))
}

// NewLogfile creates a new logger that logs to the specified file. A file is created
// with permissions specified in perm, if the file does not exist. If the file already
// exists, new records are appended to it. Additional options can be specified using opt.
func NewLogfile(file string, perm os.FileMode, opt ...Option) (log Log, err error) {
	var f io.WriteCloser
	if f, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm); err != nil {
		err = fmt.Errorf("failed to open log file; error:%w", err)
		return
	}

	return New(options(opt), WithWriter(f))
}
