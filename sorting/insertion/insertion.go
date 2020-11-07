// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 8:01 下午 - UTC/GMT+08:00

// insertion sorting
package insertion

// Sort is int type sorting func
func Sort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		pervIndex := i - 1
		current := numbers[i]
		// 用上一个元素比较当前的元素
		for pervIndex >= 0 && numbers[pervIndex] > current {
			numbers[pervIndex+1] = numbers[pervIndex]
			pervIndex -= 1
		}
		// 如果pervIndex没有变化说明就不需要操作
		if pervIndex+1 != i {
			numbers[pervIndex+1] = current
		}
	}
}

// Float64s is float64 type sorting func
func Float64s(numbers []float64) {
	for i := 1; i < len(numbers); i++ {
		pervIndex := i - 1
		current := numbers[i]
		for pervIndex >= 0 && numbers[pervIndex] > current {
			numbers[pervIndex+1] = numbers[pervIndex]
			pervIndex -= 1
		}
		numbers[pervIndex+1] = current
	}
}
