package backtracing

import (
	"fmt"
	"testing"
)

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)

	doStep(n, 0, "")

	return res
}

func doStep(open, close int, str string) {
	if open == 0 && close == 0 {
		res = append(res, str)
	}

	if open > 0 {
		doStep(open-1, close+1, str+"(")
	}

	if close > 0 {
		doStep(open, close-1, str+")")
	}
}

func Test_generateParenthesis(t *testing.T) {
	actual := generateParenthesis(1)
	fmt.Println(actual)
	actual = generateParenthesis(2)
	fmt.Println(actual)

	actual = generateParenthesis(3)
	fmt.Println(actual)
}
