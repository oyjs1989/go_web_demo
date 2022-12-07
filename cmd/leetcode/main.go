package main

import ""

func main() {
	println(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}))
	println(minOperations([]int{1, 1, 1, 1, 1, 1, 1}, []int{6}))
	println(minOperations([]int{6, 6}, []int{1}))
}
