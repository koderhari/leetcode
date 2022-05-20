package greedy

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {

		for k > 0 && len(stack) > 0 && (stack[len(stack)-1] > (byte(num[i]) - byte('0'))) {
			k--
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, byte(num[i])-byte('0'))
	}

	for k > 0 {
		k--
		stack = stack[:len(stack)-1]
	}

	var sb strings.Builder
	sb.Grow(len(stack))
	canAdd := false
	for i := 0; i < len(stack); i++ {
		if canAdd {
			sb.WriteByte(byte('0') + stack[i])
			continue
		}

		if stack[i] == 0 {
			continue
		}

		canAdd = true
		sb.WriteByte(byte('0') + stack[i])
	}

	if sb.Len() == 0 {
		return "0"
	}

	return sb.String()
}

func Test_removeKdigits(t *testing.T) {
	require.Equal(t, removeKdigits("1432219", 3), "1219")
	require.Equal(t, removeKdigits("10", 2), "0")
	require.Equal(t, removeKdigits("9", 1), "0")
}
