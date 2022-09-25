package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListItems(t *testing.T) {
	type testCase struct {
		input []string
		want  string
	}

	cases := []testCase{
		{
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
		{
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
		{
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		{
			input: []string{},
			want:  "",
		},
	}

	for _, tc := range cases {
		got := ListItems(tc.input)
		if tc.want != got {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}
