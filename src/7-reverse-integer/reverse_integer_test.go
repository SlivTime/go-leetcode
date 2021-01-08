package reverseinteger

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				123,
			},
			321,
		},
		{
			"example2",
			args{
				-123,
			},
			-321,
		},
		{
			"example3",
			args{
				120,
			},
			21,
		},
		{
			"example4",
			args{
				0,
			},
			0,
		},
		{
			"example5",
			args{
				1534236469,
			},
			0,
		},
		{
			"example6",
			args{
				-10,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
