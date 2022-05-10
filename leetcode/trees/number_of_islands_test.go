package trees

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/

var zero, one = byte('0'), byte('1')

func numIslands(grid [][]byte) int {
	res := 0
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	for i := range grid {
		for j := range grid[i] {

			if visited[i][j] {
				continue
			}

			if grid[i][j] == zero {
				continue
			}

			dfs(i, j, grid, visited)
			res++
		}
	}
	return res
}

func dfs(i int, j int, grid [][]byte, visited [][]bool) {
	n, m := len(visited), len(visited[0])

	if i < 0 || j < 0 || i >= n || j >= m {
		return
	}

	if grid[i][j] == zero {
		return
	}

	if visited[i][j] {
		return
	}

	visited[i][j] = true

	//left
	dfs(i, j-1, grid, visited)
	//right
	dfs(i, j+1, grid, visited)
	//up
	dfs(i-1, j, grid, visited)
	//dow
	dfs(i+1, j, grid, visited)
}

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: `grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]`,
			want: 1,
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
		},
		{
			name: ` [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]`,
			want: 3,
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numIslands(tc.grid)
			require.Equal(t, tc.want, actual)
		})
	}
}
