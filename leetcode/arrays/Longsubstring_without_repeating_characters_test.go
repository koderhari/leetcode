package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	res := 0
	start := 0
	visited := make(map[byte]int)
	//abcabcbbbbbbbaaaaaaacvbnmd
	for i := range s {
		idx, ok := visited[s[i]]
		if ok && start <= idx {
			c := i - start
			if c > res {
				res = c
			}

			start = idx + 1
		}

		visited[s[i]] = i
	}

	c := len(s) - start
	if c > res {
		res = c
	}

	return res
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "abcabcbb",
			in:   "abcabcbb",
			want: 3,
		},
		{
			name: "bbbbb",
			in:   "bbbbb",
			want: 1,
		},
		{
			name: "pwwkew",
			in:   "pwwkew",
			want: 3,
		},
		{
			name: "dvdf",
			in:   "dvdf",
			want: 3,
		},
		{
			name: "dvddf",
			in:   "dvddf",
			want: 2,
		},
		{
			name: "abcabcbbbbbbbaaaaaaacvbnmd",
			in:   "abcabcbbbbbbbaaaaaaacvbnmd",
			want: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(tc.in)
			require.Equal(t, tc.want, actual)
		})

	}
}
