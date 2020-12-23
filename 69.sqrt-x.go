/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (34.67%)
 * Likes:    1663
 * Dislikes: 2148
 * Total Accepted:    643K
 * Total Submissions: 1.9M
 * Testcase Example:  '4'
 *
 * Given a non-negative integer x, compute and return the square root of x.
 * 
 * Since the return type is an integer, the decimal digits are truncated, and
 * only the integer part of the result is returned.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: x = 4
 * Output: 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: x = 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since the decimal part
 * is truncated, 2 is returned.
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= x <= 2^31 - 1
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
	low, high, res := 1, x, 0

	for low <= high {
		mid := low + (high-low) >> 1
		if (mid == x/mid){
			return mid
		} else if (mid > x/mid){
			high = mid - 1
		} else {
			res = mid
			low = mid +1
		}
	}
	return res
    
}
// @lc code=end
