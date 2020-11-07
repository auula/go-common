// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/6 - 7:02 下午 - UTC/GMT+08:00

package bubble

import (
	"github.com/higker/go-common/sort"
)

//  numeric data type
//type numericType uint8
//
//const (
//	INT numericType = iota
//	INT64
//	FLOAT64
//)

var (
	positiveOrderIntFunc = func(r, l *int) {
		if *(r) > *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	reverseOrderIntFunc = func(r, l *int) {
		if *(r) < *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	positiveOrderInt64Func = func(r, l *int64) {
		if *(r) > *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	reverseOrderInt64Func = func(r, l *int64) {
		if *(r) < *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	positiveOrderFloat64Func = func(r, l *float64) {
		if *(r) > *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	reverseOrderFloat64Func = func(r, l *float64) {
		if *(r) < *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
)

// sortInt int type sort func
func sortInt(numbers []int, swap func(r, l *int)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// sortInt64 int64 type sort func
func sortInt64(numbers []int64, swap func(r, l *int64)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// sortFloat64 Float64 type sort func
func sortFloat64(numbers []float64, swap func(r, l *float64)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// Ints is int type bubble sort func
func Ints(symbol rune, numbers []int) (err error) {
	switch symbol {
	case sort.IsMoreThan:
		sortInt(numbers, positiveOrderIntFunc)
	case sort.IsLessThan:
		sortInt(numbers, reverseOrderIntFunc)
	default:
		err = sort.SymbolError
	}
	return
}

// Int64s is int64 type bubble sort func
func Int64s(symbol rune, numbers []int64) (err error) {
	switch symbol {
	case sort.IsMoreThan:
		sortInt64(numbers, positiveOrderInt64Func)
	case sort.IsLessThan:
		sortInt64(numbers, reverseOrderInt64Func)
	default:
		err = sort.SymbolError
	}
	return
}

// Float64s is float64 type bubble sort func
func Float64s(symbol rune, numbers []float64) (err error) {
	switch symbol {
	case sort.IsMoreThan:
		sortFloat64(numbers, positiveOrderFloat64Func)
	case sort.IsLessThan:
		sortFloat64(numbers, reverseOrderFloat64Func)
	default:
		err = sort.SymbolError
	}
	return
}

// golang type of data => https://studygolang.com/articles/11869
