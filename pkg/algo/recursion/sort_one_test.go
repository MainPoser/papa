package recursion

import "testing"

func Test_sortOne(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test_one", args: args{arr: []int{1, 5, 4, 7, 8, 2, 0, 9, 99, 1000, 233, 523}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortOne(tt.args.arr)
		})
	}
}
