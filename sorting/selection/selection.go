// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 1:23 下午 - UTC/GMT+08:00

package selection

import (
	"github.com/higker/go-common/sorting"
)

var (
	positiveOrderIntFunc = func(leftNum, rightNum *int) {
		if *(leftNum) > *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
	reverseOrderIntFunc = func(leftNum, rightNum *int) {
		if *(leftNum) < *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
	positiveOrderInt64Func = func(leftNum, rightNum *int64) {
		if *(leftNum) > *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
	reverseOrderInt64Func = func(leftNum, rightNum *int64) {
		if *(leftNum) < *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
	positiveOrderFloat64Func = func(leftNum, rightNum *float64) {
		if *(leftNum) > *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
	reverseOrderFloat64Func = func(leftNum, rightNum *float64) {
		if *(leftNum) < *(rightNum) {
			*(leftNum), *(rightNum) = *(rightNum), *(leftNum)
		}
	}
)

// sortInt is int type sorting func
func sortInt(numbers []int, swap func(leftNum, rightNum *int)) {
	for i := 0; i < len(numbers)-1; i++ {
		index := i
		for j := i + 1; j < len(numbers); j++ {
			swap(&numbers[index], &numbers[j])
		}
	}
}

// sortInt64 is int type sorting func
func sortInt64(numbers []int64, swap func(leftNum, rightNum *int64)) {
	for i := 0; i < len(numbers)-1; i++ {
		index := i
		for j := i + 1; j < len(numbers); j++ {
			swap(&numbers[index], &numbers[j])
		}
	}
}

// sortInt64 is int type sorting func
func sortFloat64(numbers []float64, swap func(leftNum, rightNum *float64)) {
	for i := 0; i < len(numbers)-1; i++ {
		index := i
		for j := i + 1; j < len(numbers); j++ {
			swap(&numbers[index], &numbers[j])
		}
	}
}

// Ints is int type selection sorting func
// symbol = '>' or '<'
func Ints(symbol rune, numbers []int) (err error) {
	switch symbol {
	case sorting.IsMoreThan:
		sortInt(numbers, positiveOrderIntFunc)
	case sorting.IsLessThan:
		sortInt(numbers, reverseOrderIntFunc)
	default:
		err = sorting.SymbolError
	}
	return
}

// Int64s is int64 type selection sorting func
// symbol = '>' or '<'
func Int64s(symbol rune, numbers []int64) (err error) {
	switch symbol {
	case sorting.IsMoreThan:
		sortInt64(numbers, positiveOrderInt64Func)
	case sorting.IsLessThan:
		sortInt64(numbers, reverseOrderInt64Func)
	default:
		err = sorting.SymbolError
	}
	return
}

// Float64s is float64 type selection sorting func
// symbol = '>' or '<'
func Float64s(symbol rune, numbers []float64) (err error) {
	switch symbol {
	case sorting.IsMoreThan:
		sortFloat64(numbers, positiveOrderFloat64Func)
	case sorting.IsLessThan:
		sortFloat64(numbers, reverseOrderFloat64Func)
	default:
		err = sorting.SymbolError
	}
	return
}
