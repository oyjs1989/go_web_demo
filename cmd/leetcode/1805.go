package leetcode

func numDifferentIntegers(word string) int {
	before := nil
	ret := new(map[int]string)
	for ch := range word {
		if '0' <= ch <= '9' {
			if before {
				before += ch
			} else {
				before = ch
			}
		} else {
			if before {
				ret[int(before)]
			}
			before = nil
		}
	}
	return len(ret)
}
