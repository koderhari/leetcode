package math

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func mySqrt_naive(x int) int {
	candidate := x
	for {
		candidate = candidate / 2
		v := candidate * candidate
		if x == v {
			return candidate
		}

		if v < x {
			break
		}
	}

	l, h := candidate, candidate*2

	for {
		if l > h {
			break
		}

		mid := l + (h-l)/2
		candidate = mid
		v := mid * mid
		if v == x {
			return mid
		}

		if v > x {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	if candidate*candidate > x {
		return candidate - 1
	}

	return candidate
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	l, h := 1, x
	candidate := 1
	for {
		if l > h {
			break
		}

		mid := l + (h-l)/2
		candidate = mid
		v := mid * mid
		if v == x {
			return mid
		}

		if v > x {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	if candidate*candidate > x {
		return candidate - 1
	}

	return candidate
}

func Test_mySqrt(t *testing.T) {
	//res := mySqrt(8)
	res := mySqrt(2147483647)
	require.Equal(t, 1, res)
}
