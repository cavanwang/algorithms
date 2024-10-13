package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// 定义一个整型最小堆
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
// 这是表示从小到大排序，如果<则表示从大到小排序。
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push 方法接受 interface{}，所以我们需要先类型断言
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 创建一个随机整数切片并进行堆排序
func heapSort(arr []int) {
	h := &IntHeap{}
	heap.Init(h)

	// 将切片元素添加到堆中
	for _, v := range arr {
		// 添加元素到数组末尾，然后heap会在Push内部调整堆顺序，就是v[i]往v[i/2]调整(孩子往父亲方向)
		heap.Push(h, v)
	}

	// 从堆中移除元素并添加到结果切片中
	sorted := make([]int, h.Len())
	for h.Len() > 0 {
		sorted[h.Len()-1] = heap.Pop(h).(int)
	}

	// 输出排序结果
	fmt.Println(sorted)
}

func main() {
	// 创建一个包含随机整数的切片
	rand.Seed(42)
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Int() % 100 // 生成随机整数
	}

	// 对切片进行堆排序
	heapSort(arr)
}
