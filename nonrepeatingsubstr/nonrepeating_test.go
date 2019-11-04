package nonrepeatingsubstr

import "testing"

func TestSubstr(t *testing.T)  {
	tests := []struct {
		s string
		ans int
	}{
		{"bbtablud", 6},
		{"abcbacbbab", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"b", 1},
		{"", 0},
		{"这里是慕课网", 6},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr(); " + "got %d; expected %d", actual, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8

	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer() // 前面这几行代码的时间不算，从现在开始记录benchmark运行时间

	// 循环次数b.N是系统自动计算得到的，不需要手动指定
	// 结果为 xxx ns/op，每个op即循环体内的所有语句
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("lengthOfNonRepeatingSubStr(); " + "got %d; expected %d", actual, ans)
		}
	}
}