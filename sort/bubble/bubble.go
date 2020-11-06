// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/6 - 7:02 下午 - UTC/GMT+08:00

package bubble

import (
	"errors"
)

//  numeric data type
type numericType uint8

const (
	INT numericType = iota
	INT64
	FLOAT64
)

var (
	PositiveOrderIntFunc = func(r, l *int) {
		if *(r) > *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
	ReverseOrderIntFunc = func(r, l *int) {
		if *(r) < *(l) {
			*(r), *(l) = *(l), *(r)
		}
	}
)

// AtInt int type sort func
func AtInt(numbers []int, swap func(r, l *int)) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			swap(&numbers[j], &numbers[j+1])
		}
	}
}

// AtInt64 int64 type sort func
func AtInt64(numbers []float64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

// AtInt Float64 type sort func
func AtFloat64(numbers []float64) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

// Ints is bubble int sort func
func Ints(symbol rune, numbers []int) (err error) {
	err = nil
	switch symbol {
	case '>':
		AtInt(numbers, PositiveOrderIntFunc)
	case '<':
		AtInt(numbers, ReverseOrderIntFunc)
	default:
		err = errors.New("operation symbol error, want > OR < ")
	}
	return
}
