//给定一个二叉树，找出其最大深度。 
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例： 
//给定二叉树 [3,9,20,null,null,15,7]， 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最大深度 3 。 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1181 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	//----------------------------递归-------------------------------
	// Time: O(n)
	// Space: O(height)
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left) , maxDepth(root.Right)) + 1
	//-------------------------------------------------------------

	//----------------------------迭代------------------------------
	/*// Time: O(n)
	// Space: O(n)
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0 ; i < size ; i++ {
			cur := queue[0]
			queue = queue[1 : ]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}
	return depth*/
	//-----------------------------------------------------------
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)
