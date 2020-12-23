/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (35.46%)
 * Likes:    6512
 * Dislikes: 570
 * Total Accepted:    869.7K
 * Total Submissions: 2.5M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * You are given an integer array nums sorted in ascending order, and an
 * integer target.
 * 
 * Suppose that nums is rotated at some pivot unknown to you beforehand (i.e.,
 * [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 * 
 * If target is found in the array return its index, otherwise, return -1.
 * 
 * 
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * All values of nums are unique.
 * nums is guranteed to be rotated at some pivot.
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
func search(nums []int, target int) int {
	low, high := 0, len(nums) -1
	
	for low <= high {
		mid := low + (high - low) >> 1
		if (nums[mid] == target){
			return mid
		} else if nums[mid] > nums[low]{
			if nums[low] <= target && nums[mid] > target{
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high]{
			if nums[high] >= target && nums[mid] < target{
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[mid] == nums[low]{
				low++
			}
			if (nums[mid] == nums[high]){
				high--
			}
		}
	}
	return -1
}
// @lc code=end
