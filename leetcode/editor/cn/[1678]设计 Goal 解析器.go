//请你设计一个可以解释字符串 command 的 Goal 解析器 。command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成。Goal 解
//析器会将 "G" 解释为字符串 "G"、"()" 解释为字符串 "o" ，"(al)" 解释为字符串 "al" 。然后，按原顺序将经解释得到的字符串连接成一个字
//符串。 
//
// 给你字符串 command ，返回 Goal 解析器 对 command 的解释结果。 
//
// 
//
// 示例 1： 
//
// 输入：command = "G()(al)"
//输出："Goal"
//解释：Goal 解析器解释命令的步骤如下所示：
//G -> G
//() -> o
//(al) -> al
//最后连接得到的结果是 "Goal"
// 
//
// 示例 2： 
//
// 输入：command = "G()()()()(al)"
//输出："Gooooal"
// 
//
// 示例 3： 
//
// 输入：command = "(al)G(al)()()G"
//输出："alGalooG"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= command.length <= 100 
// command 由 "G"、"()" 和/或 "(al)" 按某种顺序组成 
// 
// Related Topics 字符串 👍 25 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func interpret(command string) string {
	//------------------------------模拟-------------------------------------
	/*// Time: O(n)
	// Space: O(1)
	ans := []byte{}
	for i := 0 ; i < len(command) ; i++ {
		if command[i] == 'G' {
			ans = append(ans, 'G')
		}else if i + 1 < len(command) && command[i] == '(' && command[i + 1] == ')' {
			ans = append(ans, 'o')
		}else if i + 3 < len(command) && command[i] == '(' && command[i + 1] == 'a' && command[i + 2] == 'l' && command[i + 3] == ')' {
			ans = append(ans, 'a')
			ans = append(ans, 'l')
		}
	}
	return string(ans)*/
	//----------------------------------------------------------------------

	//------------------------------内置函数---------------------------------
	// Time: O(n)
	// Space: O(1)
	s1 := strings.ReplaceAll(command, "()", "o")
	s2 := strings.ReplaceAll(s1, "(al)", "al")
	return s2
	//----------------------------------------------------------------------
}
//leetcode submit region end(Prohibit modification and deletion)
