/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.57%)
 * Likes:    1463
 * Dislikes: 6006
 * Total Accepted:    320.5K
 * Total Submissions: 1.9M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division, and mod operator.
 * 
 * Return the quotient after dividing dividend by divisor.
 * 
 * The integer division should truncate toward zero, which means losing its
 * fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) =
 * -2.
 * 
 * Note:
 * 
 * 
 * Assume we are dealing with an environment that could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For this
 * problem, assume that your function returns 2^31 − 1 when the division result
 * overflows.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: dividend = 10, divisor = 3
 * Output: 3
 * Explanation: 10/3 = truncate(3.33333..) = 3.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: dividend = 7, divisor = -3
 * Output: -2
 * Explanation: 7/-3 = truncate(-2.33333..) = -2.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: dividend = 0, divisor = 1
 * Output: 0
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: dividend = 1, divisor = 1
 * Output: 1
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * -2^31 <= dividend, divisor <= 2^31 - 1
 * divisor != 0
 * 
 * 
 */
package leetcode


import ("math")

// @lc code=start
func divide(dividend int, divisor int) int {

	if divisor == 1 {
		return dividend
	}

	if dividend == 0 {
		return 0
	}

	sign, res := -1, 0

	if (dividend >= 0 && divisor >= 0) || (dividend < 0 && divisor < 0){
		sign = 1
	}

	low, high := 0, abs(dividend)

	for low <= high{
		quotient := low + (high-low) >> 1

		if ((quotient+1) * abs(divisor) >= abs(dividend) && quotient * abs(divisor) < abs(dividend))|| ((quotient+1) * abs(divisor) > abs(dividend) && quotient * abs(divisor) <= abs(dividend)){
			if (quotient+1) * abs(divisor) == abs(dividend){
				res = quotient+1
				break
			}
			res = quotient
			break
		} else if ((quotient+1) * abs(divisor) > abs(dividend) && quotient * abs(divisor) > abs(dividend)){
			high = quotient -1
		} else {
			low = quotient + 1
		}
	}
	
	if res > math.MaxInt32{
		return sign * math.MaxInt32
	}
	return sign * res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end
