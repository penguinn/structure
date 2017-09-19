package sort

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name:"无数测试",
			args:args{
				[]int{},
			},
			want:[]int{},
		},
		{
			name:"单个数测试",
			args:args{
				[]int{4},
			},
			want:[]int{4},
		},
		{
			name:"多个数测试",
			args:args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want:[]int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
