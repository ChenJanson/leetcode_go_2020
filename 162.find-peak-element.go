/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (43.68%)
 * Likes:    2328
 * Dislikes: 2387
 * Total Accepted:    432.5K
 * Total Submissions: 989.9K
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is strictly greater than its neighbors.
 * 
 * Given an integer array nums, find a peak element, and return its index. If
 * the array contains multiple peaks, return the index to any of the peaks.
 * 
 * You may imagine that nums[-1] = nums[n] = -∞.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 5
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2, or index number 5 where the peak element is 6.
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 1000
 * -2^31 <= nums[i] <= 2^31 - 1
 * nums[i] != nums[i + 1] for all valid i.
 * 
 * 
 * 
 * Follow up: Could you implement a solution with logarithmic complexity?
 */

// @lc code=start
func findPeakElement(nums []int) int {
	// low, high := 0, len(nums) - 1
	// for low < high {
	// 	mid := low + (high - low) >> 1
	// 	if nums[mid] > nums[mid+1]{
	// 		high = mid
	// 	} else {
	// 		low = mid + 1
	// 	}
	// }
	// return low
	if len(nums) == 0 || len(nums) == 1{
		return 0
	}
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + (high - low) >>1
		// if (mid == 0 && nums[1] < nums[0]) || (mid == len(nums)-1 && nums[mid] > nums[mid-1]) || ((mid > 0 && nums[mid] > nums[mid-1]) && (mid < len(nums)-1 && nums[mid]> nums[mid+1])){
		// 	return mid
		// }
		if (mid == len(nums)-1 && nums[mid-1] < nums[mid]) || (mid > 0 && nums[mid-1] < nums[mid] && (mid <= len(nums)-2 && nums[mid+1] < nums[mid])) || (mid == 0 && nums[1] < nums[0]) {
			return mid
		}
		if (mid > 0 && (nums[mid] > nums[mid-1])){
			low = mid +1
		} else if (mid > 0 && (nums[mid] < nums[mid-1])){
			high = mid - 1
		}
		// if mid > 0 && nums[mid-1] > nums[mid] {
		// 	high = mid - 1
		// }

		// if mid > 0 && nums[mid-1] < nums[mid] {
		// 	low = mid + 1
		// }

		if (low == mid){
			low ++
		}
		if (high == mid){
			high --
		}
	}
	return -1
	// low, high := 0, len(nums)-1
	// for low <= high {
	// 	mid := low + (high-low)>>1
	// 	if (mid == len(nums)-1 && nums[mid-1] < nums[mid]) || (mid > 0 &&
	// 		nums[mid-1] < nums[mid] && (mid <= len(nums)-2 && nums[mid+1] < nums[mid])) ||
	// 		(mid == 0 && nums[1] < nums[0]) {
	// 		return mid
	// 	}
	// 	if mid > 0 && nums[mid-1] < nums[mid] {
	// 		low = mid + 1
	// 	}
	// 	if mid > 0 && nums[mid-1] > nums[mid] {
	// 		high = mid - 1
	// 	}
	// 	if mid == low {
	// 		low++
	// 	}
	// 	if mid == high {
	// 		high--
	// 	}
	// }
	// return -1
}
// @lc code=end
