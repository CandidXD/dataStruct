package main

// BubbleSort 冒泡排序
func BubbleSort(array []int) []int {
	for j := 1; j < len(array); j++ {
		for i := 0; i < len(array)-j; i++ {
			if array[i+1] < array[i] {
				// 元素交换
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}
	return array
}
