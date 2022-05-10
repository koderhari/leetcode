package dynamic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 0; i < n; i++ {
		if (i + 1) < n {
			dp[i+1] += dp[i]
		}

		if (i + 2) < n {
			dp[i+2] += dp[i]
		}
	}

	return dp[n-1]
}

func Test_climbStairs(t *testing.T) {
	climbStairs(4)
	require.Equal(t, 5, climbStairs(4))
}

type Entity struct {
	Id int
}

var monsters []Entity
var target *Entity = nil

func TestNum(t *testing.T) {
	var monsters []*Entity
	//entity := &Entity{}
	monsters = append(monsters, &Entity{Id: 1})
	monsters = append(monsters, &Entity{Id: 2})
	monsters = append(monsters, &Entity{Id: 3})
	var target *Entity
	for i, monster := range monsters {
		if i == 2 {
			target = monster
		}
	}

	t.Log(target)
}
