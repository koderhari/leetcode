package dynamic

func maxSubArray(nums []int) int {
	res := -10_000_000
	sum := 0

	for i := range nums {
		sum = sum + nums[i]
		if sum > res {
			res = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return res
}
