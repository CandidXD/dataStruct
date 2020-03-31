package main

func QuickSort(array *[]int) {
	// 如果数组长度为1，结束
	if len(*array) <= 1 {
		return
	}
	// 除首位外排序
	// 获取左右坐标
	left := 1
	right := len(*array) - 1
	// 如果left小于right
	for left < right {
		// right找出小于等于目标数的坐标
		for right > left {
			if (*array)[right] <= (*array)[0] {
				break
			}
			right--
		}
		// left找出大于目标数的坐标
		for left < right {
			if (*array)[left] > (*array)[0] {
				break
			}
			left++
		}
		// 交换
		(*array)[right], (*array)[left] = (*array)[left], (*array)[right]
		// 坐标相同时与首位交换
	}
	if left >= right {
		// 	比较left小于首位
		if (*array)[left] <= (*array)[0] {
			// 交换
			(*array)[0], (*array)[left] = (*array)[left], (*array)[0]
		}
	}
	// 分堆
	l := (*array)[:left]
	QuickSort(&l)
	r := (*array)[right+1:]
	QuickSort(&r)
}
