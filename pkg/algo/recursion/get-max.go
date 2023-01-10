package recursion

import (
	"fmt"
)

func printMax() {
	arr := []int{1, 4, 2, 7, 4, 2, 9, 1, 89, 3}
	fmt.Printf("process(arr): %v\n", process(arr, 0, len(arr)-1))
	demo := -8
	demo = demo >> 1
	fmt.Printf("demo: %v\n", demo)
}

// process 递归获取数组中的最大值
func process(arr []int, left, right int) int {
	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)
	if left == right {
		return arr[left]
	}
	// 位运算，右移几位就是除2的几次方，左移几位就是乘上2的几次方
	// 此处这种写法是l+r的变相。防止l和r都过大导致的溢出问题的
	mid := left + ((right - left) >> 1)
	leftMax := process(arr, left, mid)
	rightMax := process(arr, mid+1, right)
	max := leftMax
	if rightMax > leftMax {
		max = rightMax
	}
	return max
}
