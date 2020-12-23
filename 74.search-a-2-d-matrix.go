/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (37.12%)
 * Likes:    2638
 * Dislikes: 185
 * Total Accepted:    398K
 * Total Submissions: 1.1M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 * 
 * 
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 3
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 13
 * Output: false
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: matrix = [], target = 0
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 0 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 * 
 * 
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// row := len(matrix)
	column := len(matrix[0])
	lowR, highR := 0, len(matrix) -1
	// lowC, highC := 0, len(matirx[0]) - 1
	for lowR <= highR{
		mid := lowR + (highR - lowR) >> 1
		if (matrix[mid][0] <= target && target <= matrix[mid][column - 1]){
			return binarySearch(matrix[mid], target)
		} else if (target < matrix[mid][0]){
			highR = mid -1
		} else if (target > matrix[mid][column - 1]){
			lowR = mid + 1
		}
	}
	return false
    
}

func binarySearch(nums []int, target int) bool {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + (high-low) >> 1
		if (nums[mid] == target){
			return true
		} else if (nums[mid] > target){
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
// @lc code=end
