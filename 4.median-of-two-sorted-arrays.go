/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (30.57%)
 * Likes:    8698
 * Dislikes: 1335
 * Total Accepted:    812.6K
 * Total Submissions: 2.7M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return
 * the median of the two sorted arrays.
 * 
 * Follow up: The overall run time complexity should be O(log (m+n)).
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums1 = [0,0], nums2 = [0,0]
 * Output: 0.00000
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: nums1 = [], nums2 = [1]
 * Output: 1.00000
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: nums1 = [2], nums2 = []
 * Output: 2.00000
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 * 
 * 
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2){
		return findMedianSortedArrays(nums2, nums1)
	} 

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1) + len(nums2) + 1)>>1, 0, 0

	for low <= high{
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid

		if (nums1Mid > 0) && (nums1[nums1Mid-1] > nums2[nums2Mid]){
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && (nums1[nums1Mid] < nums2[nums2Mid-1]){
			low = nums1Mid + 1
		} else {
			break;
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1) + len(nums2)) &1 == 1{
		return float64(midLeft)
	}

	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}

	return float64(midLeft + midRight) / 2
}

func max(a int,b int) int {
	if a > b{
		return a
	} else {
		return b
	}
}

func min(a int,b int) int {
	if a > b{
		return b
	} else {
		return a
	}
}


// @lc code=end
