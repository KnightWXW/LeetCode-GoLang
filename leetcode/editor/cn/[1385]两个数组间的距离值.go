//给你两个整数数组 arr1 ， arr2 和一个整数 d ，请你返回两个数组之间的 距离值 。 
//
// 「距离值」 定义为符合此距离要求的元素数目：对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= 
//d 。 
//
// 
//
// 示例 1： 
//
// 输入：arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
//输出：2
//解释：
//对于 arr1[0]=4 我们有：
//|4-10|=6 > d=2 
//|4-9|=5 > d=2 
//|4-1|=3 > d=2 
//|4-8|=4 > d=2 
//所以 arr1[0]=4 符合距离要求
//
//对于 arr1[1]=5 我们有：
//|5-10|=5 > d=2 
//|5-9|=4 > d=2 
//|5-1|=4 > d=2 
//|5-8|=3 > d=2
//所以 arr1[1]=5 也符合距离要求
//
//对于 arr1[2]=8 我们有：
//|8-10|=2 <= d=2
//|8-9|=1 <= d=2
//|8-1|=7 > d=2
//|8-8|=0 <= d=2
//存在距离小于等于 2 的情况，不符合距离要求 
//
//故而只有 arr1[0]=4 和 arr1[1]=5 两个符合距离要求，距离值为 2 
//
// 示例 2： 
//
// 输入：arr1 = [1,4,2,3], arr2 = [-4,-3,6,10,20,30], d = 3
//输出：2
// 
//
// 示例 3： 
//
// 输入：arr1 = [2,1,100,3], arr2 = [-5,-2,10,-3,7], d = 6
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= arr1.length, arr2.length <= 500 
// -10^3 <= arr1[i], arr2[j] <= 10^3 
// 0 <= d <= 100 
// 
// Related Topics 数组 双指针 二分查找 排序 👍 42 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	//-----------------------------模拟--------------------------------------
	/*// Time: O(m * n)
	// Space: O(1)
	ans := 0
	for i := 0 ; i < len(arr1) ; i++ {
		flag := true
		for j := 0 ; j < len(arr2) ; j++ {
			if abs(arr1[i] - arr2[j]) <= d {
				flag = false
			}
		}
		if flag == true {
			ans += 1
		}
	}
	return ans*/
	//---------------------------------------------------------------------

	//----------------------------二分查找1----------------------------------
	// Time: O(logn)
	// Space: O(1)
	sort.Ints(arr2)
	ans := 0
	for i := 0 ; i < len(arr1) ; i++ {
		// 找到 第一个 < target - d 的 元素索引：
		target, left, right := arr1[i], -1, len(arr2)
		for left + 1 != right {
			mid := left + (right - left) >> 1
			if arr2[mid] < target - d {
				left = mid
			}else{
				right = mid
			}
		}
		// 1、检查 是否全部元素 < target - d
		// 2、检查 arr2[left + 1] 是否 大于 target + d
		if (left + 1 >= len(arr2)) || (left + 1 < len(arr2) && arr2[left + 1] > target + d){
			ans++
		}
	}
	return ans
	//---------------------------------------------------------------------

	//------------------------------二分查找2---------------------------------
	// Time: O(logn)
	// Space: O(1)
	sort.Ints(arr2)
	ans := 0
	for i := 0; i < len(arr1); i++ {
		// 找到的一个 > target + d 的 元素索引：
		target, left, right := arr1[i], -1, len(arr2)
		for left + 1 != right {
			mid := left + (right - left) >> 1
			if arr2[mid] > target + d {
				right = mid
			}else{
				left = mid
			}
		}
		// 1、检查 是否全部元素 > target + d
		// 2、检查 arr2[right - 1] 是否 小于 target - d
		if (right - 1 < 0) || (right - 1 >= 0 && arr2[right - 1] < target - d) {
			ans++
		}
	}
	return ans
	//---------------------------------------------------------------------
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
