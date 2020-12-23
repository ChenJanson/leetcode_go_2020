/*
 * @lc app=leetcode id=327 lang=golang
 *
 * [327] Count of Range Sum
 *
 * https://leetcode.com/problems/count-of-range-sum/description/
 *
 * algorithms
 * Hard (35.75%)
 * Likes:    925
 * Dislikes: 102
 * Total Accepted:    46.8K
 * Total Submissions: 131K
 * Testcase Example:  '[-2,5,-1]\n-2\n2'
 *
 * Given an integer array nums, return the number of range sums that lie in
 * [lower, upper] inclusive.
 * Range sum S(i, j) is defined as the sum of the elements in nums between
 * indices i and j (i ≤ j), inclusive.
 * 
 * Note:
 * A naive algorithm of O(n^2) is trivial. You MUST do better than that.
 * 
 * Example:
 * 
 * 
 * Input: nums = [-2,5,-1], lower = -2, upper = 2,
 * Output: 3 
 * Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective
 * sums are: -2, -1, 2.
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 10^4
 * 
 * 
 */

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	res, n := 0, len(nums)
	
	for i:=0; i < n; i++ {
		temp := 0
		for j:=i; j < n;j++ {
			if i == j{
				temp = nums[j]
			} else {
				temp += nums[j]
			}

			if temp >= lower && temp <= upper {
				res ++
			}
		}
	}
	return res

	
}
// @lc code=end
