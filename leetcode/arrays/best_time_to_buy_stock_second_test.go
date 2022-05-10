package arrays

import "testing"
import "github.com/stretchr/testify/require"

//Best Time to Buy and Sell Stock II

func maxProfit(prices []int) int {
	res := 0
	buyPrices := prices[0]
	for i := 0; i < len(prices); i++ {
		if buyPrices <= prices[i] {
			res += prices[i] - buyPrices
			buyPrices = prices[i]
			continue
		}

		if buyPrices > prices[i] {
			buyPrices = prices[i]
			continue
		}
	}

	return res
}

func Test_maxProfit(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 6, 5},
			expected: 0,
		},
		{
			prices:   []int{7, 1, 5},
			expected: 4,
		},
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			actual := maxProfit(tc.prices)
			require.Equal(t, tc.expected, actual)
		})
	}
}
