//请你编写一个程序来计算两个日期之间隔了多少天。 
//
// 日期以字符串形式给出，格式为 YYYY-MM-DD，如示例所示。 
//
// 
//
// 示例 1： 
//
// 输入：date1 = "2019-06-29", date2 = "2019-06-30"
//输出：1
// 
//
// 示例 2： 
//
// 输入：date1 = "2020-01-15", date2 = "2019-12-31"
//输出：15
// 
//
// 
//
// 提示： 
//
// 
// 给定的日期是 1971 年到 2100 年之间的有效日期。 
// 
// Related Topics 数学 字符串 👍 47 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func daysBetweenDates(date1 string, date2 string) int {
	//---------------------------------模拟-------------------------------------
	// Time: O(n)
	// Space: O(1)
	arr1, arr2 := strings.Split(date1, "-"), strings.Split(date2, "-")

	y1, _ := strconv.Atoi(arr1[0])
	m1, _ := strconv.Atoi(arr1[1])
	d1, _ := strconv.Atoi(arr1[2])

	y2, _ := strconv.Atoi(arr2[0])
	m2, _ := strconv.Atoi(arr2[1])
	d2, _ := strconv.Atoi(arr2[2])

	return abs(toDays(y1, m1, d1) - toDays(y2, m2, d2))
	//--------------------------------------------------------------------------

	//-----------------------------内置time库------------------------------------
	/*// Time: O(n)
	// Space: O(1)
	time1, _ := time.Parse("2006-01-02",date1)
	time2, _ := time.Parse("2006-01-02",date2)
	dis := math.Abs(float64(time1.Unix()) - float64(time2.Unix()))
	return int(dis / (24 * 60 * 60))*/
	//--------------------------------------------------------------------------
}

func isLeapYear(y int) bool {
	if (y % 4 == 0 && y % 100 != 0) || y % 400 == 0 {
		return true
	}else{
		return false
	}
}

func toDays(year, month, day int) int {
	ans := 0

	// 年：
	for i := 1971 ; i < year ; i++ {
		if isLeapYear(i) {
			ans += 366
		}else{
			ans += 365
		}
	}

	// 月：
	arrMon := [...]int{0,31,28,31,30,31,30,31,31,30,31,30,31}
	if month > 2 && isLeapYear(year) {
		ans++
	}
	for i := 1 ; i < month ; i++ {
		ans += arrMon[i]
	}

	// 日：
	ans += day

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
//leetcode submit region end(Prohibit modification and deletion)
