package main

func InsertionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {

		insertVal := array[i+1]
		insertIndex := i + 1

		for j := i; j >= 0; j-- {
			if array[j] > insertVal {
				array[j+1] = array[j] + array[j+1]
				array[j] = array[j+1] - array[j]
				array[j+1] = array[j+1] - array[j]
				insertIndex = j
			} else {
				break
			}
		}
		array[insertIndex] = insertVal
	}
	return array
}
