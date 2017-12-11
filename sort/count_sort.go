package sort

import (
	"fmt"
)

/*插入排序*/

/*插入排序思想比较简单，分别从list[1]开始设置标兵，把标兵和前面比较，然后移位，插入标兵*/
func InsertSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}

	for i := 1; i < size; i++ {
		tmp := inArray[i]
		j := i - 1
		for ; j >= 0; j-- {
			if inArray[j] < tmp {
				break
			}
			inArray[j+1] = inArray[j]
		}
		inArray[j+1] = tmp
		fmt.Println(inArray)
	}
	return inArray
}

/*希尔排序也是一种插入排序，只是间隔值不再是1，而是从大到小*/
func ShellSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}

	blank := size / 2
	for blank > 0 {
		fmt.Println(blank)
		inArray = shellSort(inArray, size, blank)
		blank /= 2
	}
	return inArray
}

func shellSort(inArray []int, size int, blank int) []int {
	i := blank
	for ; i < size; i++ {
		tmp := inArray[i]
		j := i - blank
		for ; j >= 0; j = j - blank {
			if inArray[j] < tmp {
				break
			}
			inArray[j+blank] = inArray[j]
		}
		inArray[j+blank] = tmp
	}
	return inArray
}

/*选择排序*/

/*选择排序，选择一个最小的数放在最前面*/
func SelectSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}
	for i := 0; i < size-1; i++ {
		min := 100000000
		var index int
		for j := i; j < size; j++ {
			if inArray[j] < min {
				min = inArray[j]
				index = j
			}
		}
		inArray[i], inArray[index] = inArray[index], inArray[i]
	}
	return inArray
}

/*堆排序，建立一个小顶堆，然后每次选择根节点*/
func HeapSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}

	inArray = heapBuild(inArray, size)

	outArray := []int{}

	for j := 0; j < size-1; j++ {
		outArray = append(outArray, inArray[0])
		inArray[0] = inArray[size-1-j]
		inArray = inArray[0 : size-1-j]
		inArray = heapAdjust(inArray, 0, size-1-j)
	}
	outArray = append(outArray, inArray[0])
	return outArray
}

func heapBuild(inArray []int, size int) []int {
	for i := (size-1)/2 - 1; i >= 0; i-- {
		inArray = heapAdjust(inArray, i, size)
	}
	return inArray
}

func heapAdjust(inArray []int, i int, size int) []int {
	//有右叶子节点
	tmp := inArray[i]
	child := i*2 + 1
	for child < size-1 {
		if child+1 < size-1 && inArray[child+1] < inArray[child] {
			child++
		}
		if inArray[child] < tmp {
			inArray[i] = inArray[child]
			i = child
			child = child*2 + 1
		} else {
			break
		}
	}
	inArray[i] = tmp
	return inArray
}

func QuickSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}
	quickSort(&inArray, 0, size-1)
	return inArray
}

func quickSort(inArray *[]int, begin int, end int) {
	if begin < end {
		pos := position(inArray, begin, end)
		quickSort(inArray, begin, pos-1)
		quickSort(inArray, pos+1, end)
	}
}

func position(inArray *[]int, begin int, end int) int {
	if begin >= end {
		return begin
	}
	target := (*inArray)[end]
	storeIndex := begin
	for i := begin; i < end; i++ {
		if (*inArray)[i] < target {
			(*inArray)[storeIndex], (*inArray)[i] = (*inArray)[i], (*inArray)[storeIndex]
			storeIndex++
		}
	}

	(*inArray)[storeIndex], (*inArray)[end] = (*inArray)[end], (*inArray)[storeIndex]
	return storeIndex
}

/*基数排序*/
func CountSort(inArray []int) []int {
	size := len(inArray)
	if size == 0 {
		return inArray
	}
	//取得最大数
	var maxNum int
	for _, num := range inArray {
		if num > maxNum {
			maxNum = num
		}
	}

	//用delArray来计数
	delArray := make([]int, maxNum+1) //申明一个slice
	for _, num := range inArray {
		delArray[num]++
	}

	//用outArray来返回处理好的inArray
	outArray := make([]int, 0) //这里要用make才会有地址
	for i, count := range delArray {
		fmt.Println(outArray)
		for count > 0 {
			outArray = append(outArray, i)
			count--
		}
	}
	return outArray
}
