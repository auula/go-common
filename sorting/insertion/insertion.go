// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 8:01 下午 - UTC/GMT+08:00

package insertion

var (
	name = "1"
)

func Sort(numbers []int) {
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
