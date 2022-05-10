package arrays

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)

	for i := range strs {
		hf := hashFunc(strs[i])
		grouped[hf] = append(grouped[hf], strs[i])
	}

	res := make([][]string, 0, len(grouped))
	for i := range grouped {
		res = append(res, grouped[i])
	}

	return res
}

func hashFunc(in string) string {
	cnts := [26]byte{}
	for i := range in {
		cnts[in[i]-97]++
	}

	return string(cnts[:])
}

/*

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

tat att
a att tat bat
t tat att bat

tat
t
at
att bat
100 99 100
tat
att

e at
a te






e a t
t e a
  a n

eat
e t[,a]
*/

func Test_groupAnagrams(t *testing.T) {
	/*strs = []
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]*/
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(fmt.Sprintf("%+v", groupAnagrams(in)))

}
