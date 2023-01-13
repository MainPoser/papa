package recursion

import "fmt"

type funcType func(string) bool //代表筛选逻辑函数，可以按需实现

func filter(a []string, f funcType) []string {
	for i := range a {
		if f(a[i]) {
			return filter(append(a[:i], a[i+1:]...), f)
		}
	}
	return a
}
func deleteSpecElement() {
	newStr := filter([]string{"8", "2", "1", "4", "5"}, func(s string) bool {
		return s == "1"
	})
	fmt.Printf("newStr: %v\n", newStr)
}
