package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{1}, "1"},
		{"test2", args{2}, "2"},
		{"test3", args{3}, "fizz"},
		{"test4", args{4}, "4"},
		{"test5", args{5}, "buzz"},
		{"test6", args{6}, "fizz"},
		{"test10", args{10}, "buzz"},
		{"test15", args{15}, "fizzbuzz"},
		{"test16", args{16}, "16"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fizzbuzz(tt.args.i); got != tt.want {
				t.Errorf("Fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
