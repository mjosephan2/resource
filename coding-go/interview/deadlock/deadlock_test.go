package deadlock

import "testing"

func TestFindDeadlock(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]string
		want  bool
	}{
		{
			name: "Simple deadlock (cycle)",
			graph: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
			},
			want: true,
		},
		{
			name: "No deadlock (no cycle)",
			graph: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "D"},
			},
			want: false,
		},
		{
			name: "Multiple edges with cycle",
			graph: [][]string{
				{"A", "B"},
				{"B", "C"},
				{"C", "A"},
				{"C", "D"},
			},
			want: true,
		},
		{
			name: "Disconnected graph with one cycle",
			graph: [][]string{
				{"A", "B"},
				{"B", "A"},
				{"C", "D"},
			},
			want: true,
		},
		{
			name: "Empty graph",
			graph: [][]string{},
			want: false,
		},
		{
			name: "Single node no edge",
			graph: [][]string{
				{"A", "A"},
			},
			want: true,
		},
		{
			name: "Multiple components, no cycle",
			graph: [][]string{
				{"A", "B"},
				{"C", "D"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDeadlock(tt.graph)
			if got != tt.want {
				t.Errorf("FindDeadlock() = %v, want %v", got, tt.want)
			}
		})
	}
}