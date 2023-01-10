package recursion

import "fmt"

func sortOne(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				tmp := min
				min = arr[j]
				arr[j] = tmp
			}
		}
	}
	fmt.Printf("arr: %v\n", arr)
}
