package algorithms

import "github.com/ljun20160606/leetcode/algorithms/util"

//Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.
//
//For example, given the following matrix:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//Return 6.

func maximalRectangle(matrix [][]byte) int {
	lenY, lenX := len(matrix), len(matrix[0])
	dp := make([][]int, lenY)
	return dp[lenY-1][lenX-1]
}

// 朴素做法
// 遍历节点
// 计算包含该节点的最大面积
// 1 0 0 1 1
// 1 0 1 1 1
// 做法
// 1. 维护height[len(x)]的数组
//    记录 包含 matrix[y][0-x]节点的高度
//    height[x]存在以下情况
//    (1) height[y][x] == '0' height[x] = 0
//    (2) height[y][x] == '1' height[x] += 1
// 2. 节点matrix[y][x]向左遍历, 记录包含节点matrix[y][x]的长度，并刷新max面积
// 	  如matrix[1, 5],
//	  minHeight width square
//	  2         1	   2
//	  2         2      4
//	  1         3      3
// 最大面积为4
func naiveMaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxS int
	heights := make([]int, len(matrix[0]))
	lenY, lenX := len(matrix), len(matrix[0])
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if matrix[y][x] == '0' {
				heights[x] = 0
				continue
			}
			heights[x]++
			minHeight := heights[x]
			for k := x; k >= 0 && matrix[y][k] == '1'; k-- {
				minHeight = util.MinOfTwo(heights[k], minHeight)
				maxS = util.MaxOfTwo(maxS, (x-k+1)*minHeight)
			}
		}
	}
	return maxS
}
