package atoi

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example1",
			args{
				"42",
			},
			42,
		},
		{
			"Example2",
			args{
				"-42",
			},
			-42,
		},
		{
			"Example3",
			args{
				"4193 with words",
			},
			4193,
		},
		{
			"Example4",
			args{
				"words and 987",
			},
			0,
		},
		{
			"Example5",
			args{
				"-91283472332",
			},
			-2147483648,
		},
		{
			"-1",
			args{
				"-1",
			},
			-1,
		},
		{
			"+1",
			args{
				"+1",
			},
			1,
		},
		{
			"+-1",
			args{
				"+-1",
			},
			0,
		},
		{
			"00000-42a1234",
			args{
				"00000-42a1234",
			},
			0,
		},
		{
			"   +0 123",
			args{
				"   +0 123",
			},
			0,
		},
		{
			"9223372036854775808",
			args{
				"9223372036854775808",
			},
			2147483647,
		},
		{
			"-91283472332",
			args{
				"-91283472332",
			},
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
