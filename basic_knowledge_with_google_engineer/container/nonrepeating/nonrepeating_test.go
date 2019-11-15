package main

import "testing"

func TestSubString(t *testing.T) {
	tests := []struct{substr string; length int} {
		{"abcabcbb", 3},
		{"bbbbbb", 1},
		{"你好世界", 4},
		{"pwkefa", 6},
		{"", 0},
		{"b", 1},
		{"黑化肥挥发发灰会花灰化肥挥发会发黑会飞花黑化肥挥发发灰会花灰化肥挥发会发黑会飞花", 8},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.substr); actual != tt.length {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; except %d", tt.substr, actual, tt.length)
		}
	}
}

func BenchmarkSubstr(b *testing.B)  {
	s := "黑化肥挥发发灰会花灰化肥挥发会发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; except %d", s, actual, ans)
		}
	}
}