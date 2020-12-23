/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (41.85%)
 * Likes:    1317
 * Dislikes: 259
 * Total Accepted:    226K
 * Total Submissions: 540.1K
 * Testcase Example:  '[1,3,5]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 * 
 * Find the minimum element.
 * 
 * The array may contain duplicates.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,3,5]
 * Output: 1
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,2,0,1]
 * Output: 0
 * 
 * Note:
 * 
 * 
 * This is a follow up problem to Find Minimum in Rotated Sorted Array.
 * Would allow duplicates affect the run-time complexity? How and why?
 * 
 * 
 */

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		if (nums[low] < nums[high]){
			return nums[low]
		}
		mid := low + (high-low) >> 1
		if (nums[mid] > nums[low]){
			low = mid +1
		} else if nums[mid] == nums[low] {
			low ++
		} else {
			high = mid
		}
	}
	return nums[low]
}
// @lc code=end
