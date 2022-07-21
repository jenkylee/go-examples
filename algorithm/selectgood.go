package algorithm

func SelectGoodSort(list []int) {
	n := len(list)

	// 只需循环一半
	for i := 0; i < n/2; i++ {
		minIndex := i // 最小值下标
		maxIndex := i // 最大值下标

		for j := i + 1; j < n-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j
				continue
			}

			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		if maxIndex == i && minIndex != n-i-1 {
			// 如果最大值是开头的元素，而且最小值不是最末尾的元素
			// 先将最小的元素放在开头
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
			// 然后最小的元素放在开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == n-i-1 {
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else {
			list[i], list[minIndex] = list[minIndex], list[i]
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}
	}
}
