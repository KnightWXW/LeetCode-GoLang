//给定一个二叉树的根节点 root ，返回它的 中序 遍历。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,null,2,3]
//输出：[1,3,2]
// 
//
// 示例 2： 
//
// 
//输入：root = []
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：root = [1]
//输出：[1]
// 
//
// 示例 4： 
//
// 
//输入：root = [1,2]
//输出：[2,1]
// 
//
// 示例 5： 
//
// 
//输入：root = [1,null,2]
//输出：[1,2]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目在范围 [0, 100] 内 
// -100 <= Node.val <= 100 
// 
//
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1333 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	//-------------------------------递归-----------------------------------
	// Time: O(n)
	// Space: O(n)
	ans := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans , root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
	//---------------------------------------------------------------------

	//-------------------------------迭代----------------------------------
	/*// Time: O(n)
	// Space: O(n)
	stack , ans := []*TreeNode{}, []int{}
	for root != nil || len(stack) > 0 {
		// 每棵树的左边界节点进栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 栈顶元素出栈
		root = stack[len(stack) - 1]
		stack = stack[ : len(stack) - 1]
		// 打印栈顶元素
		ans = append(ans , root.Val)
		// 遍历到右节点
		root = root.Right
	}
	return ans*/
	//---------------------------------------------------------------------
}

//leetcode submit region end(Prohibit modification and deletion)
