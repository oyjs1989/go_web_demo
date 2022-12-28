package leetcode

//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？

func uniquePaths(m int, n int) int {
	a_m := m - 1
	a_n := n - 1
	total := a_m + a_n
	// 从 total 个中选出 a_m 个
	result := combination(a_m, total)
	return result
}

func combination(n, m int) int {
	if m/2 < n {
		n = m - n
	}
	if n == 0 {
		return 1
	}
	z := float64(m)
	for i := 1; i < n; i++ {
		z *= float64((m - i))
		z /= float64((i + 1))
	}
	return int(z)
}

func main() {
	println(uniquePaths(4, 4))
}
