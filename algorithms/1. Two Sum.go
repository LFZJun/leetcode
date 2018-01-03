package algorithms

//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	t := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := t[v]; ok {
			return []int{vv, i}
		}
		t[target-v] = i
	}
	return []int{}
}

// 前提是有序
func twoSumV0(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	left, right := 0, len(nums)-1
	var t int
	for left < right {
		t = nums[left] + nums[right]
		if t == target {
			return []int{left, right}
		} else if t < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
