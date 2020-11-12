// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/12 - 10:55 下午 - UTC/GMT+08:00

package merge

// Sort is int type merge sort
func Sort(numbers []int) []int {
	var result []int
	if len(numbers) < 2 {
		return numbers
	}
	middle := len(numbers) / 2
	left := numbers[:middle]
	right := numbers[middle:]
	return func(left, right []int) []int {
		for len(left) != 0 && len(right) != 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(right, right[0])
				right = right[1:]
			}
		}
		if len(left) != 0 {
			result = append(result, left[:]...)
		}
		if len(right) != 0 {
			result = append(result, right[:]...)
		}
		return result
	}(Sort(left), Sort(right))
}
