package log

type flags int

func (f *flags) enable(flag int, enable bool) {
	v := int(*f)
	if enable {
		v |= flag
	} else {
		v &^= flag
	}

	*f = flags(v)
}

func (f flags) isEnabled(flag int) bool {
	// no flag(s) specified; check for any flags that are ON
	if flag == 0 {
		return f > 0
	}

	// check for the specified flag(s) to be ON
	return (int(f) & flag) != 0
}
