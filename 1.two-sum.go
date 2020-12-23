/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int);
	for i:=0; i < len(nums); i++ {
		another := target - nums[i];
		if _, ok := cache[another]; ok{
			return []int{cache[another], i};
		}
		cache[nums[i]] = i;
	}
	return nil;
}
// @lc code=end
