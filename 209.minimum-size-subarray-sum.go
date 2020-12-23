/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (39.01%)
 * Likes:    3122
 * Dislikes: 129
 * Total Accepted:    316.9K
 * Total Submissions: 812.1K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of n positive integers and a positive integer s, find the
 * minimal length of a contiguous subarray of which the sum ≥ s. If there isn't
 * one, return 0 instead.
 * 
 * Example: 
 * 
 * 
 * Input: s = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: the subarray [4,3] has the minimal length under the problem
 * constraint.
 * 
 * Follow up:
 * 
 * If you have figured out the O(n) solution, try coding another solution of
 * which the time complexity is O(n log n). 
 * 
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	left, right, res, sum := 0, -1, n+1, 0
	for left < n {
		if right+1 < n && sum < s{
			right ++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left ++
		}
		if sum >= s{
			res = min(res, right-left+1)
		}
	}
	if res == n+1{
		return 0
	}
	return res
}

func min(a,b int) int {
	if a < b{
		return a
	}
	return b
}
// @lc code=end
