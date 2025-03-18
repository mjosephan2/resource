package island

import "testing"

func Test_numIslandsNoEdge(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				grid: [][]byte{
					makeByteArray([]string{"1", "1", "1", "1", "0"}),
					makeByteArray([]string{"1", "1", "0", "1", "0"}),
					makeByteArray([]string{"1", "1", "0", "0", "0"}),
					makeByteArray([]string{"0", "0", "0", "0", "0"}),
				},
			},
			want: 0,
		},
		{
			name: "example 2",
			args: args{
				grid: [][]byte{
					makeByteArray([]string{"1", "1", "0", "0", "0"}),
					makeByteArray([]string{"1", "1", "0", "0", "0"}),
					makeByteArray([]string{"0", "0", "1", "0", "0"}),
					makeByteArray([]string{"0", "0", "0", "1", "1"}),
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslandsNoEdge(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
