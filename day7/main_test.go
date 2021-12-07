package main

import (
	_ "embed"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: []string{"16", "1", "2", "0", "4", "2", "7", "1", "2", "14"},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.s); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
