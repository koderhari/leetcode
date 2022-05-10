package dynamic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func rob(nums []int) int {
	res := 0
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		if (nums[i] + dp[i]) > res {
			dp[i+1] = res
			res = nums[i] + dp[i]
		} else {
			dp[i+1] = res
		}
	}

	return res
}

func Test_rob(t *testing.T) {
	act := rob([]int{2, 7, 9, 3, 1, 1, 3})
	require.Equal(t, 15, act)

	act = rob([]int{2, 7, 9, 3, 1, 3})
	require.Equal(t, 14, act)
}
