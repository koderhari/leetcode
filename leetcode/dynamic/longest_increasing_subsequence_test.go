package dynamic

import "sort"

func lengthOfLIS(nums []int) int {

	hs := make(map[int]struct{}, len(nums))
	for i := range nums {
		hs[nums[i]] = struct{}{}
	}

	sorted := make([]int, 0, len(hs))
	for i := range hs {
		sorted = append(sorted, i)
	}

	sort.Ints(sorted)

	m := len(sorted)
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		//for j := range dp[i] {
		//	dp[i][j] = -1
		//}
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if nums[i-1] == sorted[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
