package boats

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example1",
			args{
				[]int{1, 2},
				3,
			},
			1,
		},
		{
			"Example2",
			args{
				[]int{3, 2, 2, 1},
				3,
			},
			3,
		},
		{
			"Example3",
			args{
				[]int{3, 5, 3, 4},
				5,
			},
			4,
		},
		{
			"Another one",
			args{
				[]int{3, 4, 4, 1},
				4,
			},
			3,
		},
		{
			"2-2",
			args{
				[]int{2, 2},
				6,
			},
			1,
		},
		{
			"3 not 4",
			args{
				[]int{3, 8, 7, 1, 4},
				9,
			},
			3,
		},
		{
			"2-4,5",
			args{
				[]int{2, 4},
				5,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
