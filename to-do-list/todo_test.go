package main

import (
	"testing"
)

func TestTodoList_Add(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"Single item", []string{"Buy milk"}, []string{"Buy milk"}},
		{"Multiple items", []string{"A", "B", "C"}, []string{"A", "B", "C"}},
		{"Empty input", []string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &TodoList{}

			for _, item := range tt.input {
				list.Add(item)
			}

			got := list.GetAll()

			if len(got) != len(tt.expected) {
				t.Errorf("expected %d items, got %d", len(tt.expected), len(got))
				return
			}

			for i := range got {
				if got[i].Text != tt.expected[i] {
					t.Errorf("expected item %d to be %q, got %q", i, tt.expected[i], got[i].Text)
				}
			}
		})
	}
}

func TestTodoList_Delete(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		toDelete string
		expected []string
	}{
		{"Delete Existing", []string{"A", "B", "C"}, "B", []string{"A", "C"}},
		{"Delete non existing", []string{"A", "B"}, "X", []string{"A", "B"}},
		{"Delete single", []string{"A"}, "A", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &TodoList{}
			for _, item := range tt.initial {
				list.Add(item)
			}
			list.Delete(tt.toDelete)
			got := list.GetAll()

			if len(got) != len(tt.expected) {
				t.Errorf("expected %d items, got %d", len(tt.expected), len(got))
				return
			}

			for i := range got {
				if got[i].Text != tt.expected[i] {
					t.Errorf("expected item %d to be %q, got %q", i, tt.expected[i], got[i].Text)
				}
			}

		})
	}

}
