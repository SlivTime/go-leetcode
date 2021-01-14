package longest_palindromic

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example1",
			args{
				"babad",
			},
			"bab",
		},
		{
			"example2",
			args{
				"cbbd",
			},
			"bb",
		},
		{
			"example3",
			args{
				"a",
			},
			"a",
		},
		{
			"example4",
			args{
				"ac",
			},
			"a",
		},
		{
			"banana",
			args{
				"anana",
			},
			"anana",
		},
		{
			"long one",
			args{
				"babaddtattarrattatddetartrateedredividerb",
			},
			"ddtattarrattatdd",
		},
		{
			"long one",
			args{
				"babaddtattarrattatddetartrateedredividerb",
			},
			"ddtattarrattatdd",
		},
		{
			"capitalized",
			args{
				"SQQSYYSQQS",
			},
			"SQQSYYSQQS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
