// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

// 说明：

// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// 示例 1:

// 输入: [2,2,1]
// 输出: 1
// 示例 2:

// 输入: [4,1,2,1,2]
// 输出: 4

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/single-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	// m := make(map[int]bool)
	// for i := 0; i < len(nums); i++ {
	// 	_, ok := m[nums[i]]
	// 	if ok {
	// 		m[nums[i]] = false
	// 	} else {
	// 		m[nums[i]] = true
	// 	}
	// }
	// for k, v := range m {
	// 	if v {
	// 		return k
	// 	}
	// }
	// return -1
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}
