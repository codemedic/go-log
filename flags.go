package log

type flags int

// enable-s or disables specified flag(s)
func (f *flags) enable(flag int, enable bool) {
	v := int(*f)
	if enable {
		v |= flag
	} else {
		v &^= flag
	}

	*f = flags(v)
}

// isEnabled checks if any of the specified flags are enabled.
//
// Returns true is all the specified flags are enabled. Returns true also when
// no flags are specified (check == 0) and at least one flag is enabled.
func (f flags) isEnabled(check int) bool {
	// no flag(s) specified; check for any flags that are ON
	if check == 0 {
		return f != 0
	}

	// check for the specified flag(s) to be ON
	return (int(f) & check) == check
}
