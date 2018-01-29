package origin

import "github.com/ljun20160606/leetcode/algorithms/util"

//     背包体积
//种类 \ 0 1 2 3 4 5 6
//    0 0 0 0 0 0 0 0
//    1 0
//    2 0
//    3 0
//    4 0

//1. 表示体积不超过 x 且可选前 0 种物品的情况下的最大总价值，没有物品可选，所以总价值为 0。
//2. 表示体积不超过 0 且可选前 y 种物品的情况下的最大总价值，没有物品可选，所以总价值为 0。
//3. 因为 y 这件物品的体积已经超过所能允许的最大体积了，所以肯定不能放这件物品， 那么只能在前 y-1 件物品里选了。
//4. y 这件物品可能放入背包也可能不放入背包，所以取前两者的最大值就好了， 这样就将前两种情况都包括进来了。

type item struct {
	volume int
	price  int
}

func knapsack(items []item, knapsackLoad int) int {
	lenX, lenY := knapsackLoad+1, len(items)+1
	matrix := make([][]int, lenY)
	for y := 0; y < lenY; y++ {
		matrix[y] = make([]int, lenX)
		matrix[y][0] = 0
	}
	for x := 1; x < lenX; x++ {
		matrix[0][x] = 0
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			if v := items[y-1].volume; v > x {
				matrix[y][x] = matrix[y-1][x]
			} else {
				matrix[y][x] = util.MaxOfTwo(matrix[y-1][x], items[y-1].price+matrix[y-1][x-v])
			}
		}
	}
	return matrix[len(items)][knapsackLoad]
}
