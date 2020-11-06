// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/6 - 7:02 下午 - UTC/GMT+08:00

package bubble

import (
	"errors"
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

// atInt int type sort func
func atInt(numbers []int, swap func(r, l *int)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// atInt64 int64 type sort func
func atInt64(numbers []int64, swap func(r, l *int64)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// atInt Float64 type sort func
func atFloat64(numbers []float64, swap func(r, l *float64)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// Ints is int type bubble sort func
func Ints(symbol rune, numbers []int) (err error) {
	err = nil
	switch symbol {
	case '>':
		atInt(numbers, positiveOrderIntFunc)
	case '<':
		atInt(numbers, reverseOrderIntFunc)
	default:
		err = errors.New("operation symbol error, want > OR < ")
	}
	return
}

// Int64s is int64 type bubble sort func
func Int64s(symbol rune, numbers []int64) (err error) {
	err = nil
	switch symbol {
	case '>':
		atInt64(numbers, positiveOrderInt64Func)
	case '<':
		atInt64(numbers, reverseOrderInt64Func)
	default:
		err = errors.New("operation symbol error, want > OR < ")
	}
	return
}

// Float64s is float64 type bubble sort func
func Float64s(symbol rune, numbers []float64) (err error) {
	err = nil
	switch symbol {
	case '>':
		atFloat64(numbers, positiveOrderFloat64Func)
	case '<':
		atFloat64(numbers, reverseOrderFloat64Func)
	default:
		err = errors.New("operation symbol error, want > OR < ")
	}
	return
}

// golang type of data => https://studygolang.com/articles/11869
