package bowling

import "testing"

func TestCalcScore(t *testing.T) {
	type args struct {
		bins []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{
			bins: []string{"6", "3", "9", "0", "0", "3", "8", "2", "7", "3", "X", "9", "1", "8", "0", "X", "6", "4", "5"},
		}, 139},
		{"test2", args{
			bins: []string{"6", "3", "9", "0", "0", "3", "8", "2", "7", "3", "X", "9", "1", "8", "0", "X", "X", "X", "X"},
		}, 164},
		{"test3", args{
			bins: []string{"0", "10", "1", "5", "0", "0", "0", "0", "X", "X", "X", "5", "1", "8", "1", "0", "4"},
		}, 107},
		{"test4", args{
			bins: []string{"6", "3", "9", "0", "0", "3", "8", "2", "7", "3", "X", "9", "1", "8", "0", "X", "X", "0", "0"},
		}, 134},
		{"test5", args{
			bins: []string{"6", "3", "9", "0", "0", "3", "8", "2", "7", "3", "X", "9", "1", "8", "0", "X", "X", "1", "8"},
		}, 144},
		{"test6", args{
			bins: []string{"X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X", "X"},
		}, 300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcScore(tt.args.bins); got != tt.want {
				t.Errorf("CalcScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
