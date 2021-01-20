package parensvalidation

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Example1",
			args{
				"()",
			},
			true,
		},
		{
			"Example2",
			args{
				"()[]{}",
			},
			true,
		},
		{
			"Example3",
			args{
				"(]",
			},
			false,
		},
		{
			"Example4",
			args{
				"([)]",
			},
			false,
		},
		{
			"Example5",
			args{
				"{[]}",
			},
			true,
		},
		{
			"Closing",
			args{
				"]]]",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
