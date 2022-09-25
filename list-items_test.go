package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListItems(t *testing.T) {
	t.Parallel()

	input := []string{
		"a battery",
		"a key",
		"a tourist map",
	}

	want := "You can see here a battery, a key, and a tourist map."
	got := ListItems(input)

	if want != got {
		// t.Errorf("want %q, got %q", want, got)
		t.Error(cmp.Diff(want, got))
	}
}
