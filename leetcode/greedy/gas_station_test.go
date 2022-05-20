package greedy

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func canCompleteCircuit(gas []int, cost []int) int {
	needGas := 0
	start := -1
	tank := 0
	currentTank := 0

	for i := 0; i < len(gas); i++ {
		v := gas[i] - cost[i]
		if v < 0 {
			needGas = needGas + -v
			currentTank += v
		} else {
			if start == -1 {
				start = i
				currentTank = v
			} else {
				currentTank += v
			}

			tank += v
		}

		// out previous choose was failed
		if currentTank < 0 && start > -1 {
			start = -1
		}
	}

	if tank >= needGas {
		return start
	}

	return -1
}

func Test_canCompleteCircuit(t *testing.T) {
	require.Equal(t, 3, canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	require.Equal(t, -1, canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	require.Equal(t, 5, canCompleteCircuit([]int{4, 0, 0, 0, 0, 3}, []int{1, 1, 1, 1, 1, 1}))
	require.Equal(t, 0, canCompleteCircuit([]int{4}, []int{4}))
	require.Equal(t, -1, canCompleteCircuit([]int{1}, []int{2}))
}
