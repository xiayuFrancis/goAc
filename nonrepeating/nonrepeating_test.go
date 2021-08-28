package main

import "testing"

func Test_lengthOfNonRepeatingSubStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcabcdbb", args{s: "abcabcdbb"}, 4},
		{"b", args{s: "b"}, 1},
		{"abcdefg", args{s: "abcdefg"}, 7},
		{"一二三二一", args{s: "一二三二一"}, 3},
		{"黑化肥挥发发灰会花飞化肥挥发发黑会飞花", args{s: "黑化肥挥发发灰会花飞化肥挥发发黑会飞花"}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfNonRepeatingSubStr(tt.args.s); got != tt.want {
				t.Errorf("lengthOfNonRepeatingSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lengthOfNonRepeatingSubStr1(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"黑化肥挥发发灰会花飞化肥挥发发黑会飞花", args{s: "黑化肥挥发发灰会花飞化肥挥发发黑会飞花"}, 8},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				if got := lengthOfNonRepeatingSubStr(tt.args.s); got != tt.want {
					b.Errorf("lengthOfNonRepeatingSubStr() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}
