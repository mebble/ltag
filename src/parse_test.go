package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		fixture  string
		expected Output
	}{
		{
			fixture: "../fixtures/headings.txt",
			expected: Output{
				Sections: []Section{
					{"animals", []string{"cats", "tigers", "dogs", "wolves"}, []Section{}},
					{"things", []string{"pen", "book", "table"}, []Section{}},
				},
			},
		},
		{
			fixture: "../fixtures/subheadings.txt",
			expected: Output{
				Sections: []Section{
					{"animals", []string{"elephant"}, []Section{
						{"felines", []string{"cats", "tigers"}, []Section{}},
						{"canines", []string{"dogs", "wolves"}, []Section{}},
					}},
				},
			},
		},
		{
			fixture: "../fixtures/noheadings.txt",
			expected: Output{
				Sections: []Section{
					// TODO: create separate sections when separated by newlines
					{"", []string{"one", "two", "three"}, []Section{}},
				},
			},
		},
	}

	for _, tc := range tests {
		file, err := os.Open(tc.fixture)
		if err != nil {
			t.Fatal("Unable to open test fixture")
		}
		defer file.Close()

		out := Output{Sections: []Section{}}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			out.ParseLine(line)
		}
		assert.Equal(t, tc.expected, out)
	}
}

func TestSerialise(t *testing.T) {
	tests := []struct {
		input    Output
		expected string
	}{
		{
			input: Output{
				Sections: []Section{
					{"animals", []string{"cats", "tigers", "dogs", "wolves"}, []Section{}},
					{"things", []string{"pen", "book", "table"}, []Section{}},
				},
			},
			expected: "../fixtures/headings.out.txt",
		},
		{
			input: Output{
				Sections: []Section{
					{"animals", []string{"elephant"}, []Section{
						{"felines", []string{"cats", "tigers"}, []Section{}},
						{"canines", []string{"dogs", "wolves"}, []Section{}},
					}},
				},
			},
			expected: "../fixtures/subheadings.out.txt",
		},
		{
			input: Output{
				Sections: []Section{
					// TODO: see TODO above
					{"", []string{"one", "two", "three"}, []Section{}},
				},
			},
			expected: "../fixtures/noheadings.out.txt",
		},
	}

	for _, tc := range tests {
		file, err := os.Open(tc.expected)
		if err != nil {
			t.Fatal("Unable to open expected output file")
		}
		defer file.Close()

		expected := []string{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			expected = append(expected, line)
		}

		assert.Equal(t, expected, tc.input.Serialise())
	}
}
