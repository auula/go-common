// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/11/6 - 9:00 下午 - UTC/GMT+08:00

package bubble

import "testing"

// go test -v -run=TestAtInt/greater
func TestAtInt(t *testing.T) {
	type args struct {
		numbers []int
		swap    func(r, l *int)
	}
	tests := []struct {
		name string
		args args
	}{
		{"less", args{numbers: []int{0, 9, 8, 7, 6, 4, 3, 21, 1}, swap: positiveOrderIntFunc}},
		{"greater", args{numbers: []int{0, 9, 8, 7, 6, 4, 3, 21, 1}, swap: reverseOrderIntFunc}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt)
			sortInt(tt.args.numbers, tt.args.swap)
			t.Log(tt.args.numbers)
		})
	}
}

// go test -v -run=TestSortInt
func TestSortInt(t *testing.T) {
	numbers := []int{0, 9, 8, 7, 6, 4, 3, 21, 1}
	err := Ints('>', numbers)
	if err != nil {
		t.Log(err)
	}
	t.Log(numbers)
}
