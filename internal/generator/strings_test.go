package generator

import (
	"fmt"
	"testing"
)

func TestEscapePercentChars(t *testing.T) {
	type args struct {
		s string
	}

	t.Log(fmt.Sprintf("%%"))
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "% at start",
			args: args{"% test"},
			want: "%% test",
		}, {
			name: "% at end",
			args: args{"test %"},
			want: "test %%",
		}, {
			name: "% at middle",
			args: args{"test % test"},
			want: "test %% test",
		}, {
			name: "with placeholder",
			args: args{"test {{%}} % test"},
			want: "test {{%}} %% test",
		}, {
			name: "with multiple placeholders",
			args: args{"test % {{%}} % {{%}} test"},
			want: "test %% {{%}} %% {{%}} test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapePercentChars(tt.args.s); got != tt.want {
				t.Errorf("escapePercentChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
