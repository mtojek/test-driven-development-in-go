package test_driven_development_in_go

import (
	"reflect"
	"testing"
)

func Test_divisorsOf(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"n=1", args{1}, []int{1}},
		{"n=2", args{2}, []int{1,2}},
		{"n=3", args{3}, []int{1,3}},
		{"n=4", args{4}, []int{1,2,4}},
		{"n=5", args{5}, []int{1,5}},
		{"n=24", args{24}, []int{1,2,3,4,6,8,12,24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorsOf(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divisorsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
