package main

import (
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func TestValidateAlphabet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc    string
		give    string
		wantErr string
	}{
		{
			desc:    "empty",
			wantErr: "must have at least two items",
		},
		{
			desc:    "single",
			give:    "a",
			wantErr: "must have at least two items",
		},
		{
			desc: "good",
			give: "0123456789",
		},
		{
			desc:    "dupes",
			give:    "asdffghhjjkl",
			wantErr: "alphabet has duplicates: ['f' 'h' 'j']",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			var alpha alphabet
			err := alpha.Set(tt.give)
			if len(tt.wantErr) == 0 {
				td.CmpNoError(t, err)
				return
			}

			td.CmpError(t, err)
			td.CmpContains(t, err.Error(), tt.wantErr)
		})
	}
}
