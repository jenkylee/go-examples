package algorithm

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}

	depth := length/2 - 1 // 二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i // 假定最大的位置就是在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2

		if leftchild <= length-1 && arr[rightchild] > arr[topmax] { // 防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}

		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}

	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - 1
		HeapSortMax(arr, lastLen)
		if i < lastLen {
			arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
		}
	}

	return arr
}
