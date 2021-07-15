package log

import "fmt"

func New(lf func(opt *Options) (LevelLogger, error), opts ...func(opt *Options) error) (Log, error) {
	opt := newOptions()
	for _, of := range opts {
		if err := of(opt); err != nil {
			return Log{}, fmt.Errorf("failed setting options; error:%w", err)
		}
	}

	ll, err := lf(opt)
	if err != nil {
		return Log{}, fmt.Errorf("failed to create level-logger; error:%w", err)
	}

	return Log{logger: ll}, nil
}

func NewStdLog(opts ...func(opt *Options) error) (Log, error) {
	return New(newStdLog, opts...)
}
