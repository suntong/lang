// gotests -only Double ./ > Double_test.go

package mypackage

import "testing"

func TestDouble(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Test cases
		{"3", args{3}, 6},
		{"0", args{0}, 0},
		{"-1", args{-1}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Double(tt.args.i); got != tt.want {
				t.Errorf("Double() = %v, want %v", got, tt.want)
			}
		})
	}
}
