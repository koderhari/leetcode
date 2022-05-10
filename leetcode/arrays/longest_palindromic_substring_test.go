package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func longestPalindrome_dynamic(s string) string {
	res := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for ln := 0; ln < len(s); ln++ {
		for start := 0; (start + ln) < len(s); start++ {
			if start == (start + ln) {
				dp[start][start+ln] = true
				if len(res) == 0 {
					res = s[start : start+ln+1]
				}
				continue
			}

			if s[start] != s[start+ln] {
				dp[start][start+ln] = false
				continue
			}

			if ln == 1 {
				dp[start][start+ln] = true
				c := s[start : start+ln+1]
				if len(c) > len(res) {
					res = c
				}
				continue
			}

			if dp[start+1][start+ln-1] {
				c := s[start : start+ln+1]
				if len(c) > len(res) {
					res = c
				}
			}

			dp[start][start+ln] = dp[start+1][start+ln-1]
		}
	}

	return res
}

func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		c1 := expandFromCenter(s, i, i)
		c2 := expandFromCenter(s, i, i+1)
		if len(c1) > len(res) {
			res = c1
		}
		if len(c2) > len(res) {
			res = c2
		}
	}

	return res
}

func expandFromCenter(s string, i, j int) string {
	res := ""
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			return res
		} else {
			res = s[i : j+1]
		}
		i--
		j++
	}

	return res
}

func Test_longestPalindrome(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "babad",
			in:   "babad",
			want: "bab",
		},
		{
			name: "cbbd",
			in:   "cbbd",
			want: "bb",
		},
		{
			name: "abcd",
			in:   "abcd",
			want: "a",
		},
		{
			name: "aaaa",
			in:   "aaaa",
			want: "aaaa",
		},
		{
			name: "aaaaa",
			in:   "aaaaa",
			want: "aaaaa",
		},
		{
			name: "aaaab",
			in:   "aaaab",
			want: "aaaa",
		},
		{
			name: "aaab",
			in:   "aaab",
			want: "aaa",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := longestPalindrome(tc.in)
			require.Equal(t, tc.want, actual)
		})
	}
}
