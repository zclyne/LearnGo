package basic

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {a, b, c int} {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
		{12, 35, 37},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); " + "got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

// 测试性能
func BenchmarkTriangle(b *testing.B) {
	l1, l2, l3 := 30000, 40000, 50000

	// 循环次数b.N是系统自动计算得到的，不需要手动指定
	// 结果为 xxx ns/op，每个op即循环体内的所有语句
	for i := 0; i < b.N; i++ {
		actual := calcTriangle(l1, l2)
		if actual != l3 {
			b.Errorf("calcTriangle(%d, %d); " + "got %d; expected %d", l1, l2, actual, l3)
		}
	}
}