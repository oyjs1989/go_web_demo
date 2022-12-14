package leetcode

func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	var last_word rune
	change := false
	for seq, sen := range sentence {
		// fmt.Println(last_word, sen)
		if last_word != 0 && change {
			if last_word != sen {
				return false
			}
			change = false
		}
		if sen == 32 {
			change = true
			last_word = rune(sentence[seq-1])
		}
	}
	return true
}
