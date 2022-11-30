package leetcode

/*
给你一个仅由字符 '0' 和 '1' 组成的字符串 s 。一步操作中，你可以将任一 '0' 变成 '1' ，或者将 '1' 变成 '0' 。
交替字符串 定义为：如果字符串中不存在相邻两个字符相等的情况，那么该字符串就是交替字符串。例如，字符串 "010" 是交替字符串，而字符串 "0100" 不是。
返回使 s 变成 交替字符串 所需的 最少 操作数。
*/

import "fmt"

func minOperations(s string) int {
	count := 0
	oneA := 0
	oneB := 0
	twoA := 0
	twoB := 0
	for i := 0; i < len(s); i++ {
		if count == 0 {
			if s[i] == '0' {
				oneA += 1
			} else {
				oneB += 1
			}
			count += 1
		} else {
			if s[i] == '0' {
				twoA += 1
			} else {
				twoB += 1
			}
			count = 0
		}
	}
	two := twoA + oneB
	one := twoB + oneA
	if one > two {
		return two
	}
	return one
}

func test_main() {
	fmt.Println(minOperations("0100"))
	fmt.Println(minOperations("010"))
	fmt.Println(minOperations("01"))
	fmt.Println(minOperations("1111"))
}
