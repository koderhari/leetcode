package dynamic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//up
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			//right
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func Test_uniquePaths(t *testing.T) {
	require.Equal(t, 28, uniquePaths(3, 7))
	require.Equal(t, 3, uniquePaths(3, 2))
}
