package main

import (
	"fmt"
)

// 泛型定义类型约束，确保类型支持比较操作
type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

// 泛型堆排序函数
func heapSort[T Integer](arr []T) {
	n := len(arr)
	buildMaxHeap[T](arr, n)
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify[T](arr, 0, i)
	}
}

// 初始化大顶堆
func buildMaxHeap[T Integer](arr []T, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify[T](arr, i, n)
	}
}

// 最大堆调整
func maxHeapify[T Integer](arr []T, i, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify[T](arr, largest, n)
	}
}

func main() {
	// 测试整数类型
	arrInt := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heapSort[int](arrInt)
	fmt.Println("Sorted array of integers:", arrInt)

	// 测试另一种整数类型
	arrInt64 := []int64{20, 35, -15, 7, 55, 1, -22}
	heapSort[int64](arrInt64)
	fmt.Println("Sorted array of int64:", arrInt64)
}