package leetcode

import (
	"math/rand"
	"time"
)

// quickSort 使用快速排序算法，排序整型数组
func quickSort(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	// 使用随机数优化快速排序
	rand.Seed(time.Now().Unix())
	r := rand.Intn(b-a) + a

	c := b - 1
	arr[c], arr[r] = arr[r], arr[c]

	j := a
	for i := a; i < c; i++ {
		if arr[i] < arr[c] {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	arr[j], arr[c] = arr[c], arr[j]

	quickSort(arr, a, c)
	quickSort(arr, c+1, b)
}

func minOperations(nums1 []int, nums2 []int) int {
	len1 := len(num1)
	len2 := len(num2)
	if len1 > len2 && len1 > len2*6 {
		return -1
	} else if len2 > len1 && len2 > len1*6 {
		return -1
	}
	a := 0
	b := 0
	var ret_num [len1 + len2]int
	for i := range num1 {
		a += i
		ret_num = append(ret_num, 6-i)
	}
	for i := range num2 {
		b += i
		ret_num = append(ret_num, 6-i)
	}
	c := a - b
	if c == 0 {
		return 0
	}
	quickSort(ret, 0, len1+len2)
	count := 0
	if c < 0 {
		c := -c
	}
	for i := range ret {
		count += 1
		if c-i < 0 {
			break
		}
	}
	return count

}
