package longestSubstring

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
			"abcabcbb",
			args{"abcabcbb"},
			3,
		},
		{
			"bbbbb",
			args{"bbbbb"},
			1,
		},
		{
			"pwwkew",
			args{"pwwkew"},
			3,
		},
		{
			"aab",
			args{"aab"},
			2,
		},
		{
			"dbdf",
			args{"dvdf"},
			3,
		},
		{
			"abba",
			args{"abba"},
			2,
		},
		{
			"aqua",
			args{"aqua"},
			3,
		},
		{
			"Aqua",
			args{"Aqua"},
			4,
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
