package link

import "fmt"

func next(cur, max int) int {
	cur++
	if cur > max {
		cur = cur - (max + 1)
	}
	return cur
}
func pre(cur, max int) int {
	cur--
	if cur < 0 {
		cur = cur + max + 1
	}
	return cur
}

func printRingIndex() {
	max := 7
	i := next(7, max)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("pre(i, max): %v\n", pre(i, max))
}
