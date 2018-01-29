package algorithms

import "github.com/ljun20160606/leetcode/algorithms/util"

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//Example 1:
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return findKth(nums1, nums2, l/2)
	}
	return (findKth(nums1, nums2, l/2-1) + findKth(nums1, nums2, l/2)) / 2
}

func findKth(a, b []int, k int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	if len(a) == 0 {
		return float64(b[k])
	}
	switch k {
	case 0:
		return float64(util.MinOfTwo(a[0], b[0]))
	case len(a) + len(b) - 1:
		return float64(util.MaxOfTwo(a[len(a)-1], b[len(b)-1]))
	}
	pa := util.MinOfTwo(len(a)-1, k/2)  // k/2
	pb := util.MinOfTwo(len(b)-1, k-pa) // k - k/2
	if a[pa] < b[pb] {
		return findKth(a[pa:], b[:pb], pb)
	}
	return findKth(a[:pa], b[pb:], pa)
}
