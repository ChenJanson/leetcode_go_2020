/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (36.92%)
 * Likes:    4608
 * Dislikes: 182
 * Total Accepted:    610.1K
 * Total Submissions: 1.7M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 * 
 * If target is not found in the array, return [-1, -1].
 * 
 * Follow up: Could you write an algorithm with O(log n) runtime complexity?
 * 
 * 
 * Example 1:
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 * Example 2:
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 * Example 3:
 * Input: nums = [], target = 0
 * Output: [-1,-1]
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * nums is a non-decreasing array.
 * -10^9 <= target <= 10^9
 * 
 * 
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
    return []int{searchFirstEqualNum(nums, target), searchLastEqualNum(nums,target)}
}

func searchFirstEqualNum(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	
	for low <= high {
		mid := low + (high - low) >> 1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target{
			high = mid - 1
		} else {
			if (mid == 0) || nums[mid-1] != target{
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchLastEqualNum(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	
	for low <= high {
		mid := low + (high - low) >> 1
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target{
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || nums[mid+1] != target{
				return mid
			}
			low = mid + 1
		}
	}
	return -1

}
// @lc code=end
