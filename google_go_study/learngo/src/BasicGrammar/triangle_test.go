package main

import "testing"

/**
  @author: CodeWater
  @since: 2023/4/21
  @desc: 测试代码和被测代码放一块，------basic.go
**/

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d , %d); "+"got %d ; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
