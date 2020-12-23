/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 *
 * https://leetcode.com/problems/dungeon-game/description/
 *
 * algorithms
 * Hard (32.99%)
 * Likes:    2039
 * Dislikes: 46
 * Total Accepted:    118.1K
 * Total Submissions: 357.9K
 * Testcase Example:  '[[-2,-3,3],[-5,-10,1],[10,30,-5]]'
 *
 * table.dungeon, .dungeon th, .dungeon td {
 * ⁠ border:3px solid black;
 * }
 * 
 * ⁠.dungeon th, .dungeon td {
 * ⁠   text-align: center;
 * ⁠   height: 70px;
 * ⁠   width: 70px;
 * }
 * 
 * The demons had captured the princess (P) and imprisoned her in the
 * bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid
 * out in a 2D grid. Our valiant knight (K) was initially positioned in the
 * top-left room and must fight his way through the dungeon to rescue the
 * princess.
 * 
 * The knight has an initial health point represented by a positive integer. If
 * at any point his health point drops to 0 or below, he dies immediately.
 * 
 * Some of the rooms are guarded by demons, so the knight loses health
 * (negative integers) upon entering these rooms; other rooms are either empty
 * (0's) or contain magic orbs that increase the knight's health (positive
 * integers).
 * 
 * In order to reach the princess as quickly as possible, the knight decides to
 * move only rightward or downward in each step.
 * 
 * 
 * 
 * Write a function to determine the knight's minimum initial health so that he
 * is able to rescue the princess.
 * 
 * For example, given the dungeon below, the initial health of the knight must
 * be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN ->
 * DOWN.
 * 
 * 
 * 
 * 
 * -2 (K)
 * -3
 * 3
 * 
 * 
 * -5
 * -10
 * 1
 * 
 * 
 * 10
 * 30
 * -5 (P)
 * 
 * 
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * The knight's health has no upper bound.
 * Any room can contain threats or power-ups, even the first room the knight
 * enters and the bottom-right room where the princess is imprisoned.
 * 
 * 
 */

import ("math")

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {

	if len(dungeon) == 0{
		return 0
	}
	// m, n := len(dungeon), len(dungeon[0])
	// dp := make([][]int, m)
	
	// for i:=0;  i < m; i++{
	// 	dp[i] = make([]int, n)
	// }

	// dp[m-1][n-1] = max(1-dungeon[m-1][n-1], 1)

	// for i := n-2; i>=0; i--{
	// 	dp[m-1][i] = max(dp[m-1][i+1] - dungeon[m-1][i], 1)
	// }
	// for i := m-2; i>=0; i--{
	// 	dp[i][n-1] = max(dp[i+1][n-1] - dungeon[i][n-1], 1)
	// }

	// for i := m-2; i>=0; i--{
	// 	for j := n-2; j>=0; j--{
	// 		dp[i][j] = min(max(dp[i+1][j]-dungeon[i][j], 1), max(dp[i][j+1]-dungeon[i][j], 1))
	// 	}
	// }
	// return dp[0][0]

	low, high := 1, math.MaxInt64

	for low < high {
		mid := low + (high - low) >> 1
		if canCross(dungeon, mid){
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canCross(dungeon[][] int, start int) bool {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i:=0; i < m; i++{
		dp[i] = make([]int, n)
	}

	for i:=0; i < m; i++ {
		for j:=0; j < n; j ++ {
			if i == 0 && j == 0{
				dp[i][j] = start + dungeon[0][0]
			} else {
				a, b := math.MinInt64, math.MinInt64
				if i > 0 && dp[i-1][j] > 0{
					a = dp[i-1][j] + dungeon[i][j]
				}
				if j > 0 && dp[i][j-1] > 0 {
					b = dp[i][j-1] + dungeon[i][j]
				}
				dp[i][j] = max(a, b)
			}
		}
	}
	return dp[m-1][n-1] > 0 
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b{
		return a
	}
	return b
}
// @lc code=end
