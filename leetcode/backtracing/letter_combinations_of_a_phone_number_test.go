package backtracing

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

//Letter Combinations of a Phone Number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	return step(digits, 0, "", make([]string, 0))
}

func step(digits string, i int, candidate string, res []string) []string {
	lastSymbol := i == len(digits)-1
	symbols := keys[digits[i]]
	for j := range symbols {
		if lastSymbol {
			res = append(res, candidate+string([]byte{symbols[j]}))
		} else {
			res = step(digits, i+1, candidate+string([]byte{symbols[j]}), res)
		}
	}

	return res
}

var keys = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wyxz",
}

func Test_letterCombinations(t *testing.T) {
	res := letterCombinations("23")
	sort.Strings(res)
	require.Equal(t, []string{
		"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
	}, res)

	res = letterCombinations("2")
	sort.Strings(res)
	require.Equal(t, []string{
		"a", "b", "c",
	}, res)

	res = letterCombinations("22")
	sort.Strings(res)
	require.Equal(t, []string{
		"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc",
	}, res)
}
