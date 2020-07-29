package heap

import (
	"errors"
	"fmt"
)

// 堆，优先队列
type Heap struct {
	data   []int
	length int
}

func (this *Heap) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}

func (this *Heap) Size() int {
	return this.length
}

func (this *Heap) Top() (int, error) {
	fmt.Println(this.data)
	if this.length == 0 {
		return -1, errors.New("heap has no elements")
	} else {
		return this.data[0], nil
	}
}

// 初始化大顶堆
type MaxHeap struct {
	Heap
}

func NewMaxHeap(data []int) *MaxHeap {
	length := len(data)
	if length > 1 {
		for i := length/2 - 1; i >= 0; i-- {
			data = maxHeapifyFromUp(data, i, length-1)
		}
	}

	return &MaxHeap{Heap{data: data, length: length}}
}

func (this *MaxHeap) Pop() (int, error) {
	if this.length == 0 {
		return -1, errors.New("heap has no elements")
	} else if this.length == 1 {
		result := this.data[0]
		this.length--
		this.data = this.data[1:]
		return result, nil
	} else {
		result := this.data[0]
		this.data[0], this.data[this.length-1] = this.data[this.length-1], this.data[0]
		this.data = this.data[:this.length-1]
		this.length--
		this.data = maxHeapifyFromUp(this.data, 0, this.length-1)
		return result, nil
	}
}

func (this *MaxHeap) Push(num int) {
	this.data = append(this.data, num)
	this.length++
	this.data = maxHeapifyFromDown(this.data, this.length-1, 0)
}

// 大顶堆自上向下堆化
func maxHeapifyFromUp(nums []int, start, end int) []int {
	left := 2*start + 1
	for left <= end {
		tmp := left
		right := left + 1
		if right <= end && nums[right] > nums[tmp] {
			tmp = right
		}
		if nums[tmp] > nums[start] {
			nums[tmp], nums[start] = nums[start], nums[tmp]
			start = tmp
			left = start*2 + 1
		} else {
			break
		}
	}

	return nums
}

// 大顶堆自下向上堆化
func maxHeapifyFromDown(nums []int, start, end int) []int {
	father := (start - 1) / 2
	for father >= end && father != start {
		if nums[start] > nums[father] {
			nums[father], nums[start] = nums[start], nums[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}

	return nums
}

// 初始化小顶堆
type MinHeap struct {
	Heap
}

func NewMinHeap(data []int) *MinHeap {
	length := len(data)
	if length > 1 {
		for i := length/2 - 1; i >= 0; i-- {
			data = minHeapifyFromUp(data, i, length-1)
		}
	}

	return &MinHeap{Heap{data: data, length: length}}
}

func (this *MinHeap) Pop() (int, error) {
	if this.length == 0 {
		return -1, errors.New("heap has no elements")
	} else if this.length == 1 {
		result := this.data[0]
		this.length--
		this.data = this.data[1:]
		return result, nil
	} else {
		result := this.data[0]
		this.data[0], this.data[this.length-1] = this.data[this.length-1], this.data[0]
		this.data = this.data[:this.length-1]
		this.length--
		this.data = minHeapifyFromUp(this.data, 0, this.length-1)
		return result, nil
	}
}

func (this *MinHeap) Push(num int) {
	this.data = append(this.data, num)
	this.length++
	this.data = minHeapifyFromDown(this.data, this.length-1, 0)
}

// 小顶堆自上向下堆化
func minHeapifyFromUp(nums []int, start, end int) []int {
	left := 2*start + 1
	for left <= end {
		tmp := left
		right := left + 1
		if right <= end && nums[right] < nums[tmp] {
			tmp = right
		}
		if nums[tmp] < nums[start] {
			nums[tmp], nums[start] = nums[start], nums[tmp]
			start = tmp
			left = start*2 + 1
		} else {
			break
		}
	}

	return nums
}

// 小顶堆自下向上堆化
func minHeapifyFromDown(nums []int, start, end int) []int {
	father := (start - 1) / 2
	for father >= end && father != start {
		if nums[start] < nums[father] {
			nums[father], nums[start] = nums[start], nums[father]
			start = father
			father = (start - 1) / 2
		} else {
			break
		}
	}

	return nums
}
