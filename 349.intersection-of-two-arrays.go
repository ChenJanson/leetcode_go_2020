/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (64.10%)
 * Likes:    1143
 * Dislikes: 1398
 * Total Accepted:    426.1K
 * Total Submissions: 664.6K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * Given two arrays, write a function to compute their intersection.
 * 
 * Example 1:
 * 
 * 
 * Input: nums1 = [1,2,2,1], nums2 = [2,2]
 * Output: [2]
 * 
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * Output: [9,4]
 * 
 * 
 * Note:
 * 
 * 
 * Each element in the result must be unique.
 * The result can be in any order.
 * 
 * 
 * 
 * 
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	res := []int{}
	for _, n := range(nums1){
		m[n] = true
	}

	for _, n := range(nums2){
		if m[n]{
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
// @lc code=end
