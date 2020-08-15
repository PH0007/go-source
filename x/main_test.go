package main

import "testing"

func TestGetArea(t *testing.T) {
	type args struct {
		weight int
		height int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"jim",
			args{10, 20},
			200},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArea(tt.args.weight, tt.args.height); got != tt.want {
				t.Errorf("GetArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
