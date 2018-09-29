package test_driven_development_in_go

import (
	"reflect"
	"testing"
)

func Test_factorsOf(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"n=1", args{1}, nil},
		{"n=2", args{2}, []int{2}},
		{"n=3", args{3}, []int{3}},
		{"n=4", args{4}, []int{2, 2}},
		{"n=5", args{5}, []int{5}},
		{"n=6", args{6}, []int{2, 3}},
		{"n=7", args{7}, []int{7}},
		{"n=8", args{8}, []int{2, 2, 2}},
		{"n=9", args{9}, []int{3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorsOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factorsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
