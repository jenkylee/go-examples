package main

import (
	"fmt"
	"go-examples/algorithm"
)

func main() {
	// 快速排序
	arr := []int{1, 9, 20, 39, 4, 5, 45, 8, 234, 23}
	fmt.Println(algorithm.QuickSort(arr))

	// 冒泡排序
	list := []int{5, 9, 1, 6, 14, 6, 49, 25, 4, 6, 3}
	algorithm.BubbleSort(list)
	fmt.Println(list)

	// 选择排序
	algorithm.SelectSort(list)
	fmt.Println(list)

	list0 := []int{5}

	algorithm.SelectGoodSort(list0)
	fmt.Println(list0)

	list1 := []int{5, 9}
	algorithm.SelectGoodSort(list1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	algorithm.SelectGoodSort(list2)
	fmt.Println(list2)

	algorithm.SelectGoodSort(list)
	fmt.Println(list)

	list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
	algorithm.SelectGoodSort(list4)
	fmt.Println(list4)

	// 锥排序
	arr = []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(algorithm.HeapSort(arr))
}
