package transactions

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	transitions := make([]int64, 5)
	transitions[0] = 5
	transitions[1] = 10
	transitions[2] = 1
	transitions[3] = 1
	transitions[4] = 100
	test1 := transitions

	result := make([]int64, 5)
	result[0] = 100
	result[1] = 10
	result[2] = 5
	result[3] = 1
	result[4] = 1

	type args struct {
		transitions []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Sort-transitions",
			args: args{
				transitions: test1,
			},
			want: result,
		},
	}
	for _, tt := range tests {
		if got := Sort(tt.args.transitions); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Sort() = %v, want %v", got, tt.want)
		}
	}
}
