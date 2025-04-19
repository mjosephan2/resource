package worddictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary_AddWord(t *testing.T) {
	type args struct {
		words []string
	}
	type exp struct {
		search map[string]bool
	}
	tests := []struct {
		name string
		args args
		exp  exp
	}{
		{
			args: args{
				words: []string{"car", "bus", "sand"},
			},
			exp: exp{
				search: map[string]bool{
					"c":    false,
					"bus":  true,
					"..":   false,
					".and": true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWordDictionary()
			for _, word := range tt.args.words {
				w.AddWord(word)
			}
			for word, isFound := range tt.exp.search {
				assert.Equal(t, isFound, w.Search(word), word)
			}
		})
	}
}
