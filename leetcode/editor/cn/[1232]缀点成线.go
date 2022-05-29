//给定一个数组 coordinates ，其中 coordinates[i] = [x, y] ， [x, y] 表示横坐标为 x、纵坐标为 y 的点。请你来
//判断，这些点是否在该坐标系中属于同一条直线上。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
//输出：true
// 
//
// 示例 2： 
//
// 
//
// 
//输入：coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 2 <= coordinates.length <= 1000 
// coordinates[i].length == 2 
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4 
// coordinates 中不含重复的点 
// 
// Related Topics 几何 数组 数学 👍 113 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func checkStraightLine(coordinates [][]int) bool {
	//----------------------------数学--------------------------------
	// Time: O()
	// Space: O()
	var slope func(x1, y1, x2, y2, x3, y3 int) bool
	slope = func(x1, y1, x2, y2, x3, y3 int) bool {
		return (x1- x3) * (y1 - y2) == (y1 - y3) * (x1 - x2)
	}
	for i := 0 ; i < len(coordinates) - 2 ; i++ {
		if slope(coordinates[i][0], coordinates[i][1], coordinates[i + 1][0], coordinates[i + 1][1], coordinates[i + 2][0], coordinates[i + 2][1]) == false {
			return false
		}
	}
	return true
	//---------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
