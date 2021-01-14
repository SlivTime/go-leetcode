package longestsubstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example1",
			args{
				"abcabcbb",
			},
			3,
		},
		{
			"example2",
			args{
				"bbbbb",
			},
			1,
		},
		{
			"example3",
			args{
				"pwwkew",
			},
			3,
		},
		{
			"example4",
			args{
				"",
			},
			0,
		},
		{
			"example5",
			args{
				"abcdecxyzn",
			},
			7,
		},
		{
			"aabaab!bb",
			args{
				"aabaab!bb",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
