package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example",
			args{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			"test1",
			args{[]int{2, 7, 11, 15}, 13},
			[]int{0, 2},
		},
		{
			"zero",
			args{[]int{0, 4, 3, 0}, 0},
			[]int{0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
