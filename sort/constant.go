// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 1:42 下午 - UTC/GMT+08:00

package sort

import "errors"

const (
	IsLessThan = '<'
	IsMoreThan = '>'
)

var (
	SymbolError = errors.New("operation symbol error, want > OR < ")
)
