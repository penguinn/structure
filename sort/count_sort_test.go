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
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
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

func TestInsertSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShellSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		inArray []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "无数测试",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "单个数测试",
			args: args{
				[]int{4},
			},
			want: []int{4},
		},
		{
			name: "多个数测试",
			args: args{
				[]int{2, 5, 3, 0, 2, 3, 0, 3},
			},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
		{
			name: "多多个数测试",
			args: args{
				[]int{3, 1, 5, 7, 2, 4, 9, 6, 10, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.inArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
