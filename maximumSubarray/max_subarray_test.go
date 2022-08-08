package main

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "example 2",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "example 3",
			args: args{nums: []int{-10000}},
			want: -10000,
		},
		{
			name: "example 4",
			args: args{nums: []int{-2, 1}},
			want: 1,
		},
	}
	// Testing the first soultion
	t.Run("MaxSubArray1", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := MaxSubArray1(tt.args.nums); got != tt.want {
					t.Errorf("MaxSubArray1() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("MaxSubArray2", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := MaxSubArray2(tt.args.nums); got != tt.want {
					t.Errorf("MaxSubArray2() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("MaxSubArray3", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := MaxSubArray3(tt.args.nums); got != tt.want {
					t.Errorf("MaxSubArray3() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
