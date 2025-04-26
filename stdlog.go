package log

import (
	"errors"
	"fmt"
	"io"
	stdlog "log"
	"os"
)

func formatMessage(level Level, format string, value ...interface{}) string {
	return fmt.Sprintf(level.String()+": "+format, value...)
}

type stdLogger struct {
	LevelledLogger
	PrintLevelledLogger
	StdLogSorter
	StdLogFlags
	writer io.WriteCloser
	logger *stdlog.Logger
}

func (l *stdLogger) SetWriteCloser(wc io.WriteCloser) {
	l.writer = wc
}

// Write satisfies io.Writer interface so that stdLogger can be used as writer for the standard global logger.
func (l *stdLogger) Write(p []byte) (_ int, err error) {
	level := l.SortStdlog(l.PrintLevel(), p)
	if level.IsEnabled(l.level) {
		err = l.logger.Output(4, formatMessage(level, "%s", string(p)))
	}

	return
}

// Close disables and closed the logger, freeing up resources.
func (l *stdLogger) Close() {
	l.level = Disabled
	// stop using the writer before closing it
	l.logger.SetOutput(io.Discard)
	_ = l.writer.Close()
}

func (l *stdLogger) Logf(level Level, calldepth int, format string, value ...interface{}) {
	if level.IsEnabled(l.level) {
		_ = l.logger.Output(calldepth, formatMessage(level, format, value...))
	}
}

// New creates a new logger with the specified options.
func New(opt ...Option) (log Log, err error) {
	l := &stdLogger{
		StdLogFlags: StdLogFlags{
			flags: stdlog.LstdFlags,
		},
	}

	// apply default options first
	if err = CommonOptions.Apply(l); err != nil {
		err = newConfigError(err)
		return
	}

	// apply any specified options
	for _, o := range opt {
		if err = o.Apply(l); err != nil {
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
