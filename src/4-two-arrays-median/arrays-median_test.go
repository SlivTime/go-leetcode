package median

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"example",
			args{
				[]int{1, 3},
				[]int{2},
			},
			2.0,
		},
		{
			"example2",
			args{
				[]int{1, 2},
				[]int{3, 4},
			},
			2.5,
		},
		{
			"example3",
			args{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
			},
			3.5,
		},
		{
			"example4",
			args{
				[]int{4, 4},
				[]int{4, 4},
			},
			4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
