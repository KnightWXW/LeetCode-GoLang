//图像平滑器 是大小为 3 x 3 的过滤器，用于对图像的每个单元格平滑处理，平滑处理后单元格的值为该单元格的平均灰度。 
//
// 每个单元格的 平均灰度 定义为：该单元格自身及其周围的 8 个单元格的平均值，结果需向下取整。（即，需要计算蓝色平滑器中 9 个单元格的平均值）。 
//
// 如果一个单元格周围存在单元格缺失的情况，则计算平均灰度时不考虑缺失的单元格（即，需要计算红色平滑器中 4 个单元格的平均值）。 
//
// 
//
// 给你一个表示图像灰度的 m x n 整数矩阵 img ，返回对图像的每个单元格平滑处理后的图像 。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入:img = [[1,1,1],[1,0,1],[1,1,1]]
//输出:[[0, 0, 0],[0, 0, 0], [0, 0, 0]]
//解释:
//对于点 (0,0), (0,2), (2,0), (2,2): 平均(3/4) = 平均(0.75) = 0
//对于点 (0,1), (1,0), (1,2), (2,1): 平均(5/6) = 平均(0.83333333) = 0
//对于点 (1,1): 平均(8/9) = 平均(0.88888889) = 0
// 
//
// 示例 2: 
//
// 
//输入: img = [[100,200,100],[200,50,200],[100,200,100]]
//输出: [[137,141,137],[141,138,141],[137,141,137]]
//解释:
//对于点 (0,0), (0,2), (2,0), (2,2): floor((100+200+200+50)/4) = floor(137.5) = 137
//
//对于点 (0,1), (1,0), (1,2), (2,1): floor((200+200+50+200+100+100)/6) = floor(141.
//666667) = 141
//对于点 (1,1): floor((50+200+200+200+200+100+100+100+100)/9) = floor(138.888889) =
// 138
// 
//
// 
//
// 提示: 
//
// 
// m == img.length 
// n == img[i].length 
// 1 <= m, n <= 200 
// 0 <= img[i][j] <= 255 
// 
// Related Topics 数组 矩阵 👍 142 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func imageSmoother(img [][]int) [][]int {
	//-------------------------------前缀和------------------------------
	// Time: O(m * n)
	// Space: O(m * n)
	row, col := len(img), len(img[0])
	dp := make([][]int, row + 1)
	for i := range dp {
		dp[i] = make([]int, col + 1)
	}
	// 计算前缀和数组：
	for i := 1 ; i <= row ; i++ {
		for j := 1 ; j <= col ; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1] + img[i - 1][j - 1]
		}
	}

	// 平滑处理：
	ans := make([][]int, row)
	for i := range ans{
		ans[i] = make([]int, col)
	}
	for i := 0 ; i < row ; i++ {
		for j := 0 ; j < col ; j++ {
			up := max(0 , i - 1)
			down := min(row - 1, i + 1)
			left := max(0 , j - 1)
			right := min(col - 1, j + 1)
			cnt := (down - up + 1) * (right - left + 1)
			tem := dp[down + 1][right + 1] - dp[up][right + 1] - dp[down + 1][left] + dp[up][left]
			ans[i][j] = tem / cnt
		}
	}
	return ans
	//------------------------------------------------------------------
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a , b int) int {
	if a < b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
