//有一堆石头，每块石头的重量都是正整数。 
//
// 每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下： 
//
// 
// 如果 x == y，那么两块石头都会被完全粉碎； 
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。 
// 
//
// 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。 
//
// 
//
// 示例： 
//
// 
//输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。 
//
// 
//
// 提示： 
//
// 
// 1 <= stones.length <= 30 
// 1 <= stones[i] <= 1000 
// 
// Related Topics 数组 堆（优先队列） 👍 207 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

import util "container/heap"

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i] > h[j] }

func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

func lastStoneWeight(stones []int) int {
	//--------------------------------堆-------------------------------------
	// Time: O(nlogn)
	// Space: O(n)
	if len(stones) == 1 {
		return stones[0]
	}

	h := Heap{}
	for _, v := range stones {
		h = append(h, v)
	}

	util.Init(&h)

	for len(h) >= 2 {
		first := util.Pop(&h).(int)
		second := util.Pop(&h).(int)
		if first != second {
			tem := first - second
			util.Push(&h, tem)
		}
	}
	if len(h) == 1 {
		return util.Pop(&h).(int)
	}else{
		return 0
	}
	//----------------------------------------------------------------------

	//-------------------------用排序实现堆的效果---------------------------
	/*// Time: O(nlogn)
	// Space: O(logn)
	if len(stones) == 1 {
		return stones[0]
	}
	for len(stones) >= 2 {
		sort.Ints(stones)
		if stones[len(stones) - 1] == stones[len(stones) - 2] {
			stones = stones[ : len(stones) - 2]
		}else if stones[len(stones) - 1] != stones[len(stones) - 2] {
			tem := stones[len(stones) - 1] - stones[len(stones) - 2]
			stones[len(stones) - 2] = tem
			stones = stones[ : len(stones) - 1]
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}else{
		return 0
	}*/
	//----------------------------------------------------------------------
}

//leetcode submit region end(Prohibit modification and deletion)
