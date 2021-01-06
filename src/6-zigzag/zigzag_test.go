package zigzag

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example1",
			args{
				"PAYPALISHIRING",
				3,
			},
			"PAHNAPLSIIGYIR",
		},
		{
			"example2",
			args{
				"PAYPALISHIRING",
				4,
			},
			"PINALSIGYAHRPI",
		},
		{
			"one column",
			args{
				"PAYPALISHIRING",
				14,
			},
			"PAYPALISHIRING",
		},
		{
			"one long column",
			args{
				"PAYPALISHIRING",
				50,
			},
			"PAYPALISHIRING",
		},
		{
			"single element",
			args{
				"a",
				1,
			},
			"a",
		},
		{
			"single with many rows",
			args{
				"a",
				50,
			},
			"a",
		},
		{
			"ab",
			args{
				"ab",
				1,
			},
			"ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
