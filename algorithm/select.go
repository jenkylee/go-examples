package algorithm

func SelectSort(list []int) {
	n := len(list)

	if n <= 1 {
		return
	}
    
	for i := 0; i < n-1; i++ {
		min := list[i]
		minIndex := i

		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
