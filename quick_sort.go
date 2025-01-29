package main

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	quickSort(arr[:right+1])
	quickSort(arr[left:])
}
