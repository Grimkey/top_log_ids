package main

import (
	"errors"
	"testing"

	"github.com/Grimkey/ecl_hq/ecl_heap"
)

func TestParseLine_ValidInput(t *testing.T) {
	input := "16774838: {\"id\":\"9ab7247c02044c65936a467016fff6b6\"}"
	expected := &ecl_heap.LogElement{Score: 16774838, Record: "{\"id\":\"9ab7247c02044c65936a467016fff6b6\"}"}

	actual, err := ParseLine(input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if actual == nil || *actual != *expected {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestParseLine_InvalidScore(t *testing.T) {
	input := "not-a-number: {\"id\":\"invalid\"}"
	expectedErr := "score is not a number"

	actual, err := ParseLine(input)
	if err.Error() != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}

	if actual != nil {
		t.Fatalf("expected nil, got %v", actual)
	}
}

func TestParseLine_MissingColon(t *testing.T) {
	input := "100:"
	expectedErr := "no record found"

	actual, err := ParseLine(input)
	if err == nil || err.Error() != expectedErr {
		t.Fatalf("expected error %v, got %v", expectedErr, err)
	}

	if actual != nil {
		t.Fatalf("expected nil, got %v", actual)
	}
}

func TestIdFromJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{
			input:    "{\"id\":\"9ab7247c02044c65936a467016fff6b6\"}",
			expected: "9ab7247c02044c65936a467016fff6b6",
			err:      nil,
		},
		{
			input:    "{\"score\":16774838}",
			expected: "",
			err:      errors.New("'id' not found in string"),
		},
		{
			input:    "invalid-json",
			expected: "",
			err:      errors.New("invalid character 'i' looking for beginning of value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actual, err := idFromJSON(tt.input)
			if err != nil && err.Error() != tt.err.Error() {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}

			if actual != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
