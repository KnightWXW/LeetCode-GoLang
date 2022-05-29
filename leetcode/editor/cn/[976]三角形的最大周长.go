//给定由一些正数（代表长度）组成的数组 nums ，返回 由其中三个长度组成的、面积不为零的三角形的最大周长 。如果不能形成任何面积不为零的三角形，返回 0。
// 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [2,1,2]
//输出：5
// 
//
// 示例 2： 
//
// 
//输入：nums = [1,2,1]
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 10⁴ 
// 1 <= nums[i] <= 10⁶ 
// 
// Related Topics 贪心 数组 数学 排序 👍 186 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(nums []int) int {
	//----------------------------三重循环------------------------------
	/*// Time: O(n ^ 3)
	// Space: O(1)
	ans := 0
	for i := 0 ; i < len(nums) ; i++ {
		for j := i + 1 ; j < len(nums) ; j++ {
			for k := j + 1 ; k < len(nums) ; k++ {
				arr := []int{nums[i], nums[j], nums[k]}
				if nums[i] + nums[j] + nums[k] > ans && judge(arr) {
					ans = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return ans*/
	//----------------------------------------------------------------

	//--------------------------排序 + 贪心----------------------------
	// Time: O(nlogn)
	// Space: O(logn)
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2 ; i-- {
		if nums[i] < nums[i - 1] + nums[i - 2] {
			return nums[i - 2] + nums[i - 1] + nums[i]
		}
	}
	return 0
	//----------------------------------------------------------------
}

func judge(arr []int) bool {
	sort.Ints(arr)
	if arr[2] < arr[0] + arr[1] {
		return true
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
