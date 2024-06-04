package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func min[T constraints.Float | constraints.Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Slice 自定义泛型切片
type Slice[T constraints.Float | constraints.Integer] []T

// 自定义类型作为参数
func minLen[T constraints.Integer](a, b Slice[T]) int {
	if len(a) < len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(min(1, 2))
	fmt.Println(minLen([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(minLen([]int32{1, 2}, []int32{1, 2, 3}))

	fmt.Println(min[float64](1.1, 2.2))
}
