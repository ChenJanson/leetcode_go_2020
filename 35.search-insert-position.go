/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (42.72%)
 * Likes:    2960
 * Dislikes: 280
 * Total Accepted:    733.9K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array of distinct integers and a target value, return the
 * index if the target is found. If not, return the index where it would be if
 * it were inserted in order.
 * 
 * 
 * Example 1:
 * Input: nums = [1,3,5,6], target = 5
 * Output: 2
 * Example 2:
 * Input: nums = [1,3,5,6], target = 2
 * Output: 1
 * Example 3:
 * Input: nums = [1,3,5,6], target = 7
 * Output: 4
 * Example 4:
 * Input: nums = [1,3,5,6], target = 0
 * Output: 0
 * Example 5:
 * Input: nums = [1], target = 0
 * Output: 0
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums contains distinct values sorted in ascending order.
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums) - 1

	if target < nums[low]{
		return 0
	}

	if target > nums[high]{
		return high + 1
	}
	
	for low <= high {
		mid := low + (high- low) >> 1
		if (nums[mid] == target){
			return mid
		} else if (nums[mid] < target){
			low = mid + 1
		} else {
			if nums[mid-1] < target{
				return mid
			}
			high = mid - 1
		}
		
	}
	return -1
}
// @lc code=end
