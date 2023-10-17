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
	}

	for _, tc := range tests {
		out := Output{Sections: []Section{}}
		file, err := os.Open(tc.fixture)
		if err != nil {
			t.Fatal("Unable to open test fixture")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			out.ParseLine(line)
		}
		assert.Equal(t, tc.expected, out)
	}
}
