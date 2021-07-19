package log

import (
	"reflect"
	"testing"
)

func Test_options_append(t *testing.T) {
	type args struct {
		opt []Option
	}
	tests := []struct {
		name  string
		o     options
		args  args
		after options
	}{
		{
			name: "",
			o:    OptionMust(Options(WithLevel(Debug))).(options),
			args: args{
				opt: []Option{WithLevel(Info)},
			},
			after: OptionMust(Options(
				WithLevel(Debug),
				WithLevel(Info),
			)).(options),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.o.append(tt.args.opt...); !reflect.DeepEqual(tt.o, tt.after) {
				t.Errorf("append() = %v, after %v", tt.o, tt.after)
			}
		})
	}
}
