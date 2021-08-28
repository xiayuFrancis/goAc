package triangle

import "testing"

func Test_calcTriangle(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3", args{3, 4}, 5},
		{"5", args{5, 12}, 13},
		{"8", args{8, 15}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTriangle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("calcTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
