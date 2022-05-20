package greedy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; ; {
		if j < i {
			break
		}

		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestMaxArea(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	require.Equal(t, 1, maxArea([]int{2, 1}))
	require.Equal(t, 1, maxArea([]int{1, 1}))
	require.Equal(t, 1, maxArea([]int{1, 2}))

	require.Equal(t, 8, maxArea([]int{1, 8, 8, 1}))
	require.Equal(t, 24, maxArea([]int{8, 1, 1, 8}))
	require.Equal(t, 16, maxArea([]int{8, 1, 8, 1}))
	require.Equal(t, 16, maxArea([]int{1, 8, 1, 8}))

	require.Equal(t, 50, maxArea([]int{2, 1, 50, 51, 1, 2}))

	require.Equal(t, 24, maxArea([]int{1, 3, 2, 5, 25, 24, 5}))

}
