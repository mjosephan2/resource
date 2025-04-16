package ascendingindex

import (
	"reflect"
	"testing"
)

func Test_ascendingIndex(t *testing.T) {
	type args struct {
		inputList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				inputList: []int{73, 74, 75, 60, 61, 62, 80},
			},
			want: []int{1, 1, 4, 1, 1, 1, 0},
		},
		{
			args: args{
				inputList: []int{72, 72, 72},
			},
			want: []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ascendingIndex(tt.args.inputList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ascendingIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
