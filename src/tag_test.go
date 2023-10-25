package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagTransform(t *testing.T) {
	// the tests kinda follow this pattern where each successive test is more constrained by the prevous
	// so if the last test passes, then the previous tests should be covered as well. But not sure if this holds for every edge case.
	// but this is a general rule. Doing it this way because it helps to implement just the simpler tests first
	tests := []struct {
		inputFile  string
		expectedFile string
	}{
		{
			inputFile: "../fixtures/headings.txt",
			expectedFile: "../fixtures/headings.out.txt",
		},
		{
			inputFile: "../fixtures/subheadings.txt",
			expectedFile: "../fixtures/subheadings.out.txt",
		},
		{
			inputFile: "../fixtures/noheadings.txt",
			expectedFile: "../fixtures/noheadings.out.txt",
		},
		{
			inputFile: "../fixtures/inline.txt",
			expectedFile: "../fixtures/inline.out.txt",
		},
		{
			inputFile: "../fixtures/slug.txt",
			expectedFile: "../fixtures/slug.out.txt",
		},
		{
			inputFile: "../fixtures/multiheadings.txt",
			expectedFile: "../fixtures/multiheadings.out.txt",
		},
	}

	for _, tc := range tests {
		inputFile, err := os.Open(tc.inputFile)
		if err != nil {
			t.Fatal("Unable to open input file")
		}
		defer inputFile.Close()

		expectedFile, err := os.Open(tc.expectedFile)
		if err != nil {
			t.Fatal("Unable to open expected file")
		}
		defer expectedFile.Close()

		s := NewTaggingBuf("#", "#")

		inputScanner := bufio.NewScanner(inputFile)
		expectedScanner := bufio.NewScanner(expectedFile)
		for inputScanner.Scan() {
			line := inputScanner.Text()
			taggedLine, ok := s.Transform(line)
			if !ok {
				assert.Empty(t, taggedLine)
				continue
			}

			expectedScanner.Scan()
			expectedLine := expectedScanner.Text()
			assert.Equal(t, expectedLine, taggedLine)
		}
	}
}
