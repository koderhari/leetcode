package sorting

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func search(nums []int, target int) int {
	l, h, mid := 0, len(nums)-1, 0
	for {
		if l > h {
			break
		}

		mid = l + (h-l)/2

		//check mid

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= target && target < nums[mid] {
			h = mid - 1
		} else if nums[mid] < target && target <= nums[h] {
			l = mid + 1
		} else if nums[h] >= nums[mid] {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func Test_search(t *testing.T) {
	actual := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	require.Equal(t, 4, actual)

	actual = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	require.Equal(t, -1, actual)

	actual = search([]int{1}, 0)
	require.Equal(t, -1, actual)

	actual = search([]int{4, 5, 6, 7, 0, 1, 2}, 4)
	require.Equal(t, 0, actual)

	actual = search([]int{5, 6, 0, 1, 2, 3, 4}, 5)
	require.Equal(t, 0, actual)
	actual = search([]int{5, 6, 0, 1, 2, 3, 4}, 10)
	require.Equal(t, -1, actual)

	actual = search([]int{5, 6, 0, 1, 2, 3, 4}, 4)
	require.Equal(t, 6, actual)

	actual = search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
	require.Equal(t, 4, actual)
}

/*
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Input: nums = [1], target = 0
Output: -1

*/
