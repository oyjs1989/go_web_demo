package leetcode

//字节面试题，求大佬们看看，数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n的最大数。
//例如A={1, 2, 4, 9}，x=2533，返回2499
//

func maxNumber(numList []int, target int) int {
	for num := string(target){
		last := 0
		for _, n := range numList{
			if n <= num{
				last = n
			}
		}
	}
}

func main() {
	println(maxNumber({1, 2, 4, 9}, 2533))
}
