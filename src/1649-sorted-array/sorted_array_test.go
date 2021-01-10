package sortedarray

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func readFromFile(path string) []int {
	var result []int
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Test_createSortedArray(t *testing.T) {
	type args struct {
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				[]int{1, 5, 6, 2},
			},
			1,
		},
		{
			"example2",
			args{
				[]int{1, 2, 3, 6, 5, 4},
			},
			3,
		},
		{
			"example3",
			args{
				[]int{1, 3, 3, 3, 2, 4, 2, 1, 2},
			},
			4,
		},
		{
			"small one",
			args{
				[]int{1},
			},
			0,
		},
		{
			"one number",
			args{
				[]int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			0,
		},
		{
			"already sorted",
			args{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			0,
		},
		{
			"more numbers",
			args{
				[]int{4, 14, 10, 2, 5, 3, 8, 19, 7, 20, 12, 1, 9, 15, 13, 11, 18, 6, 16, 17},
			},
			43,
		},
		{
			"huge amount",
			args{
				readFromFile("many_numbers.json"),
			},
			188426454,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createSortedArray(tt.args.instructions); got != tt.want {
				t.Errorf("createSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
