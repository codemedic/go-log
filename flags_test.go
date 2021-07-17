package log

import (
	stdlog "log"
	"testing"
)

func Test_flags_enable(t *testing.T) {
	type args struct {
		flag   int
		enable bool
	}
	tests := []struct {
		name  string
		flags flags
		args  args
		after flags
	}{
		{
			name:  "disable a flag that is currently enabled",
			flags: stdlog.LUTC | stdlog.LstdFlags,
			args: args{
				flag:   stdlog.LUTC,
				enable: false,
			},
			after: stdlog.LstdFlags,
		},
		{
			name:  "disable a flag that is currently disabled",
			flags: stdlog.LstdFlags,
			args: args{
				flag:   stdlog.LUTC,
				enable: false,
			},
			after: stdlog.LstdFlags,
		},
		{
			name:  "enable a flag that is currently enabled",
			flags: stdlog.LUTC | stdlog.LstdFlags,
			args: args{
				flag:   stdlog.LUTC,
				enable: true,
			},
			after: stdlog.LUTC | stdlog.LstdFlags,
		},
		{
			name:  "enable a flag that is currently disabled",
			flags: stdlog.LstdFlags,
			args: args{
				flag:   stdlog.LUTC,
				enable: true,
			},
			after: stdlog.LUTC | stdlog.LstdFlags,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.flags.enable(tt.args.flag, tt.args.enable)
			if tt.args.enable {
				if !tt.flags.isEnabled(tt.args.flag) {
					t.Error("failed to enable", tt.args.flag)
				}
			} else {
				if tt.flags.isEnabled(tt.args.flag) {
					t.Error("failed to disable", tt.args.flag)
				}
			}

			if tt.flags != tt.after {
				t.Error("flags does not match 'after'")
			}
		})
	}
}
