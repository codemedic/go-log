package log

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

func levelFromString(str string) (Level, error) {
	switch str {
	case "error", "ERROR", "Error":
		return Error, nil
	case "warning", "WARNING", "Warning":
		return Warning, nil
	case "info", "INFO", "Info":
		return Info, nil
	case "debug", "DEBUG", "Debug":
		return Debug, nil
	case "disabled", "DISABLED", "Disabled", "":
		return Disabled, nil
	}

	return 0, ErrBadLevel
}
