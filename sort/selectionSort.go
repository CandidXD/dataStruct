package main

// SelectionSort 选择排序
func SelectionSort(array []int) []int {
	var position int
	for i := 0; i < len(array)-1; i++ {
		position = i
		for j := i; j < len(array); j++ {
			if array[position] > array[j] {
				position = j
			}
		}
		if position == i {
			continue
		}
		array[i] = array[i] + array[position]
		array[position] = array[i] - array[position]
		array[i] = array[i] - array[position]
	}

	return array
}
