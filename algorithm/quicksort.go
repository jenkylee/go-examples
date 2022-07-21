package algorithm

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	splitdata := arr[0]
	low := make([]int, 0)
	hight := make([]int, 0)
	mid := make([]int, 0)

	mid = append(mid, splitdata)

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i])
		} else if arr[i] > splitdata {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)

	myarr := append(append(low, mid...), hight...)

	return myarr
}
