//给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。 
//
// 请你找到并返回这个整数 
//
// 
//
// 示例： 
//
// 
//输入：arr = [1,2,2,6,6,6,6,7,10]
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr.length <= 10^4 
// 0 <= arr[i] <= 10^5 
// 
// Related Topics 数组 👍 55 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findSpecialInteger(arr []int) int {
	//---------------------------hashMap-------------------------------
	/*// Time: O(n)
	// Space: O(n)
	hashMap := map[int]int{}
	for _, v := range arr {
		hashMap[v]++
		if hashMap[v] * 4 > len(arr) {
			return v
		}
	}
	return -1*/
	//----------------------------------------------------------------

	//-----------------------------排序--------------------------------
	// Time: O(nlogn)
	// Space: O(logn)
	n := len(arr)
	k := n >> 2
	for i := 0 ; i < n - k ; i++ {
		if arr[i] == arr[i + k]	{
			return arr[i]
		}
	}
	return -1
	//----------------------------------------------------------------

	//------------------------------贪心-------------------------------
	// Time: O(n)
	// Space: O(1)
	ans, cnt := arr[0], 1
	for i := 0 ; i < len(arr) - 1 ; i++ {
		if arr[i] == arr[i + 1] {
			cnt++
		}else{
			cnt = 1
			ans = arr[i]
		}
		if cnt * 4 > len(arr) {
			return ans
		}
	}
	return -1
	//-----------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
