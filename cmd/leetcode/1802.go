package main

import (
	"math"
	"sort"
)

// 给你三个正整数 n、index 和 maxSum 。
// 你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：
//
// nums.length == n
// nums[i] 是 正整数 ，其中 0 <= i < n
// abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
// nums 中所有元素之和不超过 maxSum
// nums[index] 的值被 最大化
// 返回你所构造的数组中的 nums[index] 。
//
//注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。
// 示例 1：
//
//输入：n = 4, index = 2,  maxSum = 6
//输出：2
//解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。

// 示例 2：
// 输入：n = 6, index = 1,  maxSum = 10
// 输出：3
func maxValue(n int, index int, maxSum int) int {
	sum := func(x, cnt int) int {
		if x >= cnt {
			return (x + x - cnt + 1) * cnt / 2
		}
		return (x+1)*x/2 + cnt - x
	}
	return sort.Search(maxSum, func(x int) bool {
		x++
		return sum(x-1, index)+sum(x, n-index) > maxSum
	})
}

func maxValue1(n, index, maxSum int) int {
	left := index
	right := n - index - 1
	if left > right {
		left, right = right, left
	}

	upper := ((left+1)*(left+1)-3*(left+1))/2 + left + 1 + (left + 1) + ((left+1)*(left+1)-3*(left+1))/2 + right + 1
	if upper >= maxSum {
		a := 1.0
		b := -2.0
		c := float64(left + right + 2 - maxSum)
		return int((-b + math.Sqrt(b*b-4*a*c)) / (2 * a))
	}

	upper = (2*(right+1)-left-1)*left/2 + (right + 1) + ((right+1)*(right+1)-3*(right+1))/2 + right + 1
	if upper >= maxSum {
		a := 1.0 / 2
		b := float64(left) + 1 - 3.0/2
		c := float64(right + 1 + (-left-1)*left/2.0 - maxSum)
		return int((-b + math.Sqrt(b*b-4*a*c)) / (2 * a))
	} else {
		a := float64(left + right + 1)
		b := float64(-left*left-left-right*right-right)/2 - float64(maxSum)
		return int(-b / a)
	}
}

func main() {
	println(maxValue(4, 2, 6))
}
