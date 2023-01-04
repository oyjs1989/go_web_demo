package leetcode

import "strconv"

// 句子是由若干 token 组成的一个列表，token 间用 单个 空格分隔，句子没有前导或尾随空格。
// 每个 token 要么是一个由数字 0-9 组成的不含前导零的 正整数 ，要么是一个由小写英文字母组成的 单词 。

// 示例，"a puppy has 2 eyes 4 legs" 是一个由 7 个 token 组成的句子："2" 和 "4" 是数字，其他像 "puppy" 这样的 tokens 属于单词。
// 给你一个表示句子的字符串 s ，你需要检查 s 中的 全部 数字是否从左到右严格递增（即，除了最后一个数字，s 中的 每个 数字都严格小于它 右侧 的数字）。

func isNumber(s string) bool {
	//println("input string", s)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		//println("isNumber ch", ch)
		if 48 > ch || ch > 57 {
			return false
		}
	}
	return true
}

func areNumbersAscending(s string) bool {
	tmp := ""
	big := -1
	for i := 0; i < len(s); i++ {
		ch := s[i]
		//println(string(ch))
		if ch != ' ' {
			tmp += string(ch)
		} else {
			//println("tmp:", tmp)
			if isNumber(tmp) {
				now, _ := strconv.Atoi(tmp)
				//println("int:", now)
				if now > big {
					big = now
				} else {
					return false
				}
			}
			tmp = ""
		}
	}
	//println("end tmp:", tmp)
	if isNumber(tmp) {
		now, _ := strconv.Atoi(tmp)
		//println("int:", now)
		if now > big {
			big = now
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	println(areNumbersAscending("hello world 5 x 5"))
	println(areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"))
	println(areNumbersAscending("4 5 11 26"))
}
