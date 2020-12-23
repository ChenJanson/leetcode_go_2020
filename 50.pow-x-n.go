/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (30.73%)
 * Likes:    1964
 * Dislikes: 3475
 * Total Accepted:    567.6K
 * Total Submissions: 1.8M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (i.e. x^n).
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: x = 2.00000, n = 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: x = 2.10000, n = 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: x = 2.00000, n = -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * -10^4 <= x^n <= 10^4
 * 
 * 
 */

// @lc code=start
func myPow(x float64, n int) float64 {

    if n == 1 {
		return x
	}

	if n == 0 {
		return 1
	}


	if n < 0 {
		n = -n
		x = 1/x
	}

	tmp := myPow(x, n/2)

	if n % 2 == 0{
		return tmp * tmp
	}

	return tmp * tmp * x 
}
// @lc code=end
