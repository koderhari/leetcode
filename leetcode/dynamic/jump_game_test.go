package dynamic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//[2,3,1,1,4]
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := range nums {
		if !dp[i] {
			continue
		}

		if (i + nums[i]) >= len(nums) {
			return true
		}

		for j := 0; j <= nums[i]; j++ {
			dp[i+j] = true
		}
	}

	return dp[len(nums)-1]
}

func Test_canJump(t *testing.T) {
	w := canJump([]int{2, 3, 1, 1, 4})
	require.Equal(t, true, w)
}
