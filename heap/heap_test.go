package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewMaxHeap([]int{})
	heap.Push(40)
	heap.Push(12)
	heap.Push(16)
	heap.Push(14)
	heap.Push(35)
	heap.Push(19)
	heap.Push(34)
	heap.Push(35)
	heap.Push(28)
	heap.Push(35)
	heap.Push(26)
	heap.Push(6)
	heap.Push(8)
	heap.Push(2)
	heap.Push(14)
	heap.Push(25)
	heap.Push(25)
	heap.Push(4)
	heap.Push(33)
	heap.Push(18)
	heap.Push(10)
	heap.Push(14)
	fmt.Println(heap.Top())
}
