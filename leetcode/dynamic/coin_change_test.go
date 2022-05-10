package dynamic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		coin := coins[i]
		if coin > amount {
			continue
		}
		dp[coin] = 1
		for j := 1; j < len(dp); j++ {
			if j < coin && dp[j] == 0 {
				continue
			}
			if j+coin >= len(dp) {
				break
			}

			if dp[j] == 0 {
				continue
			}

			cnt := dp[j] + 1
			if dp[j+coin] == 0 || dp[j+coin] > cnt {
				dp[j+coin] = cnt
				continue
			}
		}

	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

func Test_coinChange(t *testing.T) {
	testcases := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "1 2 5",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "1 1 1",
			coins:  []int{1, 1, 1},
			amount: 11,
			want:   11,
		},
		{
			name:   "1 2 5",
			coins:  []int{1, 2, 5},
			amount: 13,
			want:   4,
		},
		{
			name:   "2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "5 2 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "[186,419,83,408]",
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			want:   20,
		},
	}
	/*Input:
	[186,419,83,408]
	6249
	Output:
	-1
	Expected:
	20*/

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := coinChange(tc.coins, tc.amount)
			require.Equal(t, tc.want, actual)
		})
	}
}
