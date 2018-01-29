package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms/util"
)

//Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
//
//For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
//the contiguous subarray [4,-1,2,1] has the largest sum = 6.

// dp[i]为包含nums[i]的最大subarray和
// dp[i]有两种情况
// 1. dp[i-1]+nums[i]
// 2. nums[i]
// dp[i]每次只与dp[i-1]有关可以优化为dp
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, dp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp = util.MaxOfTwo(dp+nums[i], nums[i])
		max = util.MaxOfTwo(dp, max)
	}
	return max
}
