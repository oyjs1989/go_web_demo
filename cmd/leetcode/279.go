package leetcode

import "math"

// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
//
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
// https://zh.wikipedia.org/zh-sg/%E5%9B%9B%E5%B9%B3%E6%96%B9%E5%92%8C%E5%AE%9A%E7%90%86
// 任何正整数都可以拆分成不超过4个数的平方和 ---> 答案只可能是1,2,3,4
// 如果一个数最少可以拆成4个数的平方和，则这个数还满足 n = (4^a)*(8b+7) ---> 因此可以先看这个数是否满足上述公式，如果不满足，答案就是1,2,3了
// 如果这个数本来就是某个数的平方，那么答案就是1，否则答案就只剩2,3了
// 如果答案是2，即n=a^2+b^2，那么我们可以枚举a，来验证，如果验证通过则答案是2
// 只能是3
func numSquares1(n int) int {
	// 动态规划
	sqrtNum := int(math.Sqrt(float64(n)))
	println(sqrtNum)
	last := n - sqrtNum*sqrtNum
	if last == 0 {
		return 1
	} else {
		return 1 + numSquares(last)
	}
}

func numSquares(n int) int {
	// 动态规划
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(numSquares(13))
}
