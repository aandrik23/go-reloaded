package main

import (
	"testing"
)

func TestModifyText(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			"1E (hex) files were added",
			"30 files were added",
		},
		{
			"It has been 10 (bin) years",
			"It has been 2 years",
		},
	}

	for _, test := range tests {
		result := modifyText(test.input)
		if result != test.output {
			t.Errorf("Expected '%s' but go '%s'", test.output, result)
		}
	}
}
