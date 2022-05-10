package trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Given the root of a binary tree, return the inorder traversal of its nodes' values.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	var stack []*TreeNode
	stack = append(stack, root)
	for {
		if len(stack) == 0 {
			break
		}

		fromStack := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if fromStack.Right == nil && fromStack.Left == nil {
			res = append(res, fromStack.Val)
			continue
		}

		if fromStack.Right != nil {
			stack = append(stack, fromStack.Right)
		}

		stack = append(stack, &TreeNode{
			Val: fromStack.Val,
		})

		if fromStack.Left != nil {
			stack = append(stack, fromStack.Left)
		}
	}

	return res
}

func Test_inorderTraversal(t *testing.T) {
	testCases := []struct {
		name string
		root func() *TreeNode
		want []int
	}{
		{
			name: "[1,null,2,3]",
			want: []int{1, 3, 2},
			root: func() *TreeNode {
				r := &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				}

				right := &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: nil,
				}

				r.Right = right

				return r
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := inorderTraversal(tc.root())
			require.Equal(t, tc.want, actual)
		})
	}
}
