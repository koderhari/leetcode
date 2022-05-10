package sorting

func sortColors(nums []int) {
	zeros, ones, twos := 0, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			zeros++
		} else if nums[i] == 1 {
			ones++
		} else {
			twos++
		}
	}

	idx := 0
	for ; zeros > 0; zeros-- {
		nums[idx] = 0
		idx++
	}

	for ; ones > 0; ones-- {
		nums[idx] = 1
		idx++
	}

	for ; twos > 0; twos-- {
		nums[idx] = 2
		idx++
	}
}
