package sorting

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//[4,]
	l, h, mid := 0, len(nums)-1, 0
	for {
		if h < l {
			break
		}

		mid = l + (h-l)/2
		if nums[mid] == target {
			break
		}

		if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	if nums[mid] != target {
		return []int{-1, -1}
	}

	start, end, i, j := mid, mid, mid, mid

	for {
		if i <= 0 && j >= len(nums)-1 {
			break
		}

		if i-1 >= 0 && nums[i-1] == target {
			start--
		}
		i--

		if j+1 < len(nums) && nums[j+1] == target {
			end++
		}

		j++
	}

	//for i := mid; i >= 0; i-- {
	//	if nums[i] == target {
	//		start = i
	//	} else {
	//		break
	//	}
	//}
	//
	//for i := mid; i < len(nums); i++ {
	//	if nums[i] == target {
	//		end = i
	//	} else {
	//		break
	//	}
	//}

	return []int{start, end}
}

func Test_searchRange(t *testing.T) {
	require.Equal(t, []int{0, 0}, searchRange([]int{1}, 1))
	require.Equal(t, []int{0, 1}, searchRange([]int{1, 1}, 1))
	require.Equal(t, []int{0, 2}, searchRange([]int{1, 1, 1}, 1))
	require.Equal(t, []int{-1, -1}, searchRange([]int{1, 1, 1}, 2))
	require.Equal(t, []int{3, 4}, searchRange([]int{1, 2, 3, 4, 4, 5}, 4))
	require.Equal(t, []int{1, 2}, searchRange([]int{1, 2, 2, 3, 4, 5}, 2))
	require.Equal(t, []int{1, 1}, searchRange([]int{1, 2, 3, 4, 5}, 2))
	require.Equal(t, []int{3, 3}, searchRange([]int{1, 2, 3, 4, 5}, 4))
}
