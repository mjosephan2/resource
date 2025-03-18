package island

import "testing"

func Test_numIslands(t *testing.T) {
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
			want: 1,
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
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeByteArray(runes []string) []byte {
	var bytes []byte
	for _, r := range runes {
		bytes = append(bytes, byte(r[0]))
	}
	return bytes
}
