package log

import "strings"

type Level int

const (
	Disabled Level = iota
	Debug
	Info
	Warning
	Error
)

func (l Level) String() string {
	switch l {
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Disabled:
		return "DISABLED"
	default:
		return "UNKNOWN"
	}
}

// IsEnabled return true if l is above or at the same level as threshold.
// If threshold is Disabled or is below l then returns false.
func (l Level) IsEnabled(threshold Level) bool {
	if threshold == Disabled {
		return false
	}

	return l >= threshold
}

func (l *Level) UnmarshalText(text []byte) error {
	parsedLevel, err := levelFromString(string(text))
	if err != nil {
		return err
	}
	*l = parsedLevel
	return nil
}

func levelFromString(str string) (Level, error) {
	switch strings.ToLower(str) {
	case "error":
		return Error, nil
	case "warning":
		return Warning, nil
	case "info":
		return Info, nil
	case "debug":
		return Debug, nil
	case "disabled", "":
		return Disabled, nil
	}
	return 0, ErrBadLevel
}
