package main

import (
	"container/heap"
	"fmt"
)

type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(a, b int) bool {
	return pq[a] < pq[b]
}

func (pq priorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq *priorityQueue) Push(val interface{}) {
	item := val.(int)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	data := *pq
	lastIndex := len(data) - 1
	item := data[lastIndex]
	*pq = data[:lastIndex]
	return item
}

func kthLargest(nums []int, k int) int {
	queue := make(priorityQueue, 0, k)
	heap.Init(&queue)

	for i := 0; i < len(nums); i++ {
		heap.Push(&queue, nums[i])

		if len(queue) > k {
			heap.Pop(&queue)
		}
	}

	return queue[0]
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	k := 2

	result := kthLargest(arr, k)
	fmt.Println("K-th largest:", result)
}
