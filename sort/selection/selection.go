// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 1:23 下午 - UTC/GMT+08:00

package selection

import (
	"github.com/higker/go-common/sort"
)

var (
	positiveOrderIntFunc = func(leftNum, rightNum *int) {
		if *(leftNum) > *(rightNum) {
			*(leftNum), *(leftNum) = *(rightNum), *(rightNum)
		}
	}
	reverseOrderIntFunc = func(leftNum, rightNum *int) {
		if *(leftNum) < *(rightNum) {
			*(leftNum), *(leftNum) = *(rightNum), *(rightNum)
		}
	}
)

// atInt is int type sorting func
func atInt(numbers []int, swap func(leftNum, rightNum *int)) {
	for i := 0; i < len(numbers)-1; i++ {
		index := i
		for j := i + 1; j < len(numbers); j++ {
			swap(&numbers[index], &numbers[j])
		}
	}
}

func Ints(symbol rune, numbers []int) (err error) {
	switch symbol {
	case sort.IsMoreThan:
		atInt(numbers, positiveOrderIntFunc)
	case sort.IsLessThan:
		atInt(numbers, reverseOrderIntFunc)
	default:
		err = sort.SymbolError
	}
	return
}
