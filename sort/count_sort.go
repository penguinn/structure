package sort

import "fmt"

func CountSort(inArray []int) []int {
	size := len(inArray)
	if size == 0{
		return  inArray
	}
	//取得最大数
	var maxNum int
	for _, num := range inArray{
		if num > maxNum{
			maxNum = num
		}
	}

	//用delArray来计数
	delArray := make([]int, maxNum+1)      	//申明一个slice
	for _, num := range inArray{
		delArray[num] ++
	}

	//用outArray来返回处理好的inArray
	outArray := make([]int, 0)					//这里要用make才会有地址
	for i, count := range delArray{
		fmt.Println(outArray)
		for count > 0{
			outArray = append(outArray, i)
			count --
		}
	}
	return outArray
}
