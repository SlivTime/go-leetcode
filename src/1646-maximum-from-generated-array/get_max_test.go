package getmax

import "testing"

func Test_getMaximumGenerated(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example1",
			args{
				7,
			},
			3,
		},
		{
			"Example2",
			args{
				2,
			},
			1,
		},
		{
			"Example3",
			args{
				3,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGenerated(tt.args.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
