//在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直
//到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？ 
//
// 
//
// 示例 1: 
//
// 输入: 
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 12
//解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物 
//
// 
//
// 提示： 
//
// 
// 0 < grid.length <= 200 
// 0 < grid[0].length <= 200 
// 
// Related Topics 数组 动态规划 矩阵 👍 281 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxValue(grid [][]int) int {
	//---------------------------递归(超时)---------------------------------
	/*// Time: O(n * m)
	// Space: O(n * m)
	var valueDFS func(grid [][]int, i, j int) int

	valueDFS = func(grid [][]int, i, j int) int {
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		if i == 0 {
			return valueDFS(grid, i, j - 1)	+ grid[i][j]
		}
		if j == 0 {
			return valueDFS(grid, i - 1, j)	+ grid[i][j]
		}
		p1, p2 := valueDFS(grid, i - 1, j), valueDFS(grid, i, j - 1)
		return max(p1, p2) + grid[i][j]
	}

	return valueDFS(grid, len(grid) - 1, len(grid[0]) - 1)*/
	//---------------------------------------------------------------

	//-------------------------记忆化搜索------------------------------
	/*// Time: O(n * m)
	// Space: O(n * m)
	mapArr := make([][]int, len(grid))
	for i := range mapArr {
		mapArr[i] = make([]int, len(grid[0]))
	}

	var valueDFS func(grid [][]int, i, j int) int

	valueDFS = func(grid [][]int, i, j int) int {
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		if i == 0 {
			if mapArr[i][j - 1] == 0 {
				return valueDFS(grid, i, j - 1)	+ grid[i][j]
			}else{
				return mapArr[i][j - 1]	+ grid[i][j]
			}
		}
		if j == 0 {
			if mapArr[i - 1][j] == 0 {
				return valueDFS(grid, i - 1, j)	+ grid[i][j]
			}else{
				return mapArr[i - 1][j]	+ grid[i][j]
			}
		}
		p1, p2 := 0, 0
		if mapArr[i][j - 1] != 0 {
			p1 = mapArr[i][j - 1]
		}else{
			p1 = valueDFS(grid, i, j - 1)
		}
		if mapArr[i - 1][j] != 0 {
			p2 = mapArr[i - 1][j]
		}else{
			p2 = valueDFS(grid, i - 1, j)
		}
		mapArr[i][j] = max(p1, p2) + grid[i][j]
		return max(p1, p2) + grid[i][j]
	}

	return valueDFS(grid, len(grid) - 1, len(grid[0]) - 1)*/
	//---------------------------------------------------------------

	//-------------------------动态规划------------------------------
	// Time: O(n * m)
	// Space: O(1)
	row, col := len(grid), len(grid[0])

	for i := 1 ; i < row ; i++ {
		grid[i][0] += grid[i - 1][0]
	}
	for j := 1 ; j < col ; j++ {
		grid[0][j] += grid[0][j - 1]
	}
	for i := 1 ; i < row ; i++ {
		for j := 1 ; j < col ; j++ {
			grid[i][j] += max(grid[i][j - 1], grid[i - 1][j])
		}
	}
	return grid[row - 1][col - 1]
	//---------------------------------------------------------------
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
