// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/12 - 10:29 下午 - UTC/GMT+08:00

package shell

// Sort is int type shell sorting
func Sort(numbers []int) {
	var pervIndex, current int
	for gap := len(numbers) / 2; gap > 0; gap /= 2 {
		for j := gap; j < len(numbers); j++ {
			current = numbers[j]
			pervIndex = j - gap
			for pervIndex >= 0 && current < numbers[pervIndex] {
				numbers[pervIndex+gap] = numbers[pervIndex]
				// 互换位置得到原来的位置
				pervIndex -= gap
			}
			numbers[pervIndex+gap] = current
		}
	}
}

// Float64s is float64 type  shell sorting func
func Float64s(numbers []float64) {
	var pervIndex, current float64
	for gap := float64(len(numbers) / 2); gap > 0; gap /= 2 {
		for j := gap; j < float64(len(numbers)); j++ {
			current = numbers[int(j)]
			pervIndex = j - gap
			for pervIndex >= 0 && current < numbers[int(pervIndex)] {
				numbers[int(pervIndex+gap)] = numbers[int(pervIndex)]
				// 互换位置得到原来的位置
				pervIndex -= gap
			}
			numbers[int(pervIndex+gap)] = current
		}
	}
}
