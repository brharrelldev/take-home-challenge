package treesearch

import (
	"testing"
)

func TestInsert(t *testing.T) {

	type pair struct {
		key  int
		data []byte
	}

	tests := []struct {
		name     string
		inputs   []pair
		expected int
	}{
		{
			name: "happy path",
			inputs: []pair{
				{1, []byte("test")},
				{2, []byte("test")},
				{5, []byte("test")},
				{8, []byte("test")},
				{3, []byte("test")},
				{4, []byte("test")},
				{5, []byte("test")},
			},
			expected: 3,
		},
		{
			name: "no duplications",
			inputs: []pair{
				{1, []byte("test")},
				{2, []byte("test")},
				{5, []byte("test")},
				{8, []byte("test")},
				{3, []byte("test")},
				{4, []byte("test")},
				{7, []byte("test")},
			},
			expected: 0,
		},
		{
			name:     "nil tree",
			inputs:   nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bst := &BinarySearchTree{}
			for _, input := range tt.inputs {
				bst.Insert(input.key, input.data)
			}

			level := CheckDuplicateId(bst.Root)

			if level != tt.expected {
				t.Fatalf("expected=%v got=%v", tt.expected, level)
			}

		})
	}
}
