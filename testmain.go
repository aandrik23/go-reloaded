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
		{
			"Ready, set, go (up) !",
			"Ready, set, GO !",
		},
		{
			"I should stop SHOUTING (low)",
			"I should stop shouting",
		},
		{
			"Welcome to the Brooklyn bridge (cap)",
			"Welcome to the Brooklyn Bridge",
		},
		{
			"This is so exciting (up, 2)",
			"This is SO EXCITING",
		},
		{
			"I was sitting over there ,and then BAMM !!",
			"I was sitting over there, and then BAMM!!",
		},
		{
			"I was thinking ... You were right",
			"I was thinking... You were right",
		},
		{
			"I am exactly how they describe me: ' awesome '",
			"I am exactly how they describe me: 'awesome'",
		},
		{
			"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			"There it was. A amazing rock!",
			"There it was. An amazing rock!",
		},
	}

	for _, test := range tests {
		result := ModifyText(test.input)
		if result != test.output {
			t.Errorf("Expected '%s' but go '%s'", test.output, result)
		}
	}
}
