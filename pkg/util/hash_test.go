package util

import (
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"
	"testing"
)

func TestDeepHashObject(t *testing.T) {
	type args struct {
		hasher        hash.Hash
		objectToWrite interface{}
	}
	testHash := fnv.New32a()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "chinese",
			args: struct {
				hasher        hash.Hash
				objectToWrite interface{}
			}{
				hasher:        testHash,
				objectToWrite: "短Has测试，啊啊啊啊啊啊哒哒哒 阿达啊啊  啊啊的a",
			},
		},
		{
			name: "number",
			args: struct {
				hasher        hash.Hash
				objectToWrite interface{}
			}{
				hasher:        testHash,
				objectToWrite: "13246548684616461616168",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeepHashObject(tt.args.hasher, tt.args.objectToWrite)
		})
	}
}

func TestComputeHash(t *testing.T) {
	type args struct {
		template       interface{}
		collisionCount *int32
	}
	i := int32(1)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "chinese",
			args: struct {
				template       interface{}
				collisionCount *int32
			}{
				template:       "短Has测试，啊啊啊啊啊啊哒哒哒 阿达啊啊  啊啊的a",
				collisionCount: &i,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeHash(tt.args.template, tt.args.collisionCount)
			t.Logf("ComputeHash() = %v", got)
		})
	}
}

func BenchmarkComputeHash(b *testing.B) {
	baseStr := "短Has测试，啊啊啊啊啊啊哒哒哒 阿达啊啊  啊啊的a"
	for i := 0; i < b.N; i++ {
		j := int32(i)
		got := ComputeHash(baseStr+strconv.Itoa(i), &j)
		fmt.Printf("times: %vgot: %v\n", j, got)
	}
}
