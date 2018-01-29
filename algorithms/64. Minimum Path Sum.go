package algorithms

import "github.com/ljun20160606/leetcode/algorithms/util"

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example 1:
//[[1,3,1],
// [1,5,1],
// [4,2,1]]
//Given the above grid map, return 7. Because the path 1→3→1→1→1 minimizes the sum.

func minPathSum(grid [][]int) int {
	lenX, lenY := len(grid[0]), len(grid)
	matrix := make([][]int, lenY)
	for sum, y := 0, 0; y < lenY; y++ {
		matrix[y] = make([]int, lenX)
		sum += grid[y][0]
		matrix[y][0] = sum
	}
	for sum, x := 0, 0; x < lenX; x++ {
		sum += grid[0][x]
		matrix[0][x] = sum
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			matrix[y][x] = util.MinOfTwo(matrix[y][x-1], matrix[y-1][x]) + grid[y][x]
		}
	}
	return matrix[lenY-1][lenX-1]
}
