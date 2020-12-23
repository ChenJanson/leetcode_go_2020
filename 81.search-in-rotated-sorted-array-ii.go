/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (33.41%)
 * Likes:    1802
 * Dislikes: 525
 * Total Accepted:    281.6K
 * Total Submissions: 842.8K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 * 
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 * 
 * Follow up:
 * 
 * 
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 * 
 * 
 */

// @lc code=start
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	
	for low <= high {
		mid := low + (high-low)>>1

		if (nums[mid] == target){
			return true
		} else if (nums[mid] > nums[low]){
			if (target >= nums[low] && target < nums[mid]){
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if (nums[mid] < nums[high]){
			if (target > nums[mid] && target <= nums[high]){
				low = mid + 1
			} else {
				high = mid -1
			} 
		}else {
				if nums[low] == nums[mid] {
					low ++
				} 
				if nums[mid] == nums[high]{
					high --
				}
			}
	}
	return false
}
// @lc code=end
