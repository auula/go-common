// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/7 - 8:45 下午 - UTC/GMT+08:00

package insertion

import "testing"

func TestSort(t *testing.T) {
	numbers := []int{2, 1, 3, 4, 98, 11, 13, 12, 9, 10}
	t.Log("UnSorted :", numbers)
	Sort(numbers)
	t.Log("Sorted :", numbers)
}
