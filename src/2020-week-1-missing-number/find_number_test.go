package find_number

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				[]int{2, 3, 4, 7, 11},
				5,
			},
			9,
		},
		{
			"example2",
			args{
				[]int{1, 2, 3, 4},
				2,
			},
			6,
		},
		{
			"single element",
			args{
				[]int{1},
				2,
			},
			3,
		},
		{
			"first one",
			args{
				[]int{2, 3, 4},
				1,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
