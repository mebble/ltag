package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/mebble/ltag/internal/ltag"
	"github.com/stretchr/testify/assert"
)

func TestTagTransform(t *testing.T) {
	// the tests kinda follow this pattern where each successive test is more constrained by the prevous
	// so if the last test passes, then the previous tests should be covered as well. But not sure if this holds for every edge case.
	// but this is a general rule. Doing it this way because it helps to implement just the simpler tests first
	tests := []struct {
		inputFile    string
		expectedFile string
	}{
		{
			inputFile:    "./testdata/headings.txt",
			expectedFile: "./testdata/headings.out.txt",
		},
		{
			inputFile:    "./testdata/subheadings.txt",
			expectedFile: "./testdata/subheadings.out.txt",
		},
		{
			inputFile:    "./testdata/noheadings.txt",
			expectedFile: "./testdata/noheadings.out.txt",
		},
		{
			inputFile:    "./testdata/inline.txt",
			expectedFile: "./testdata/inline.out.txt",
		},
		{
			inputFile:    "./testdata/slug.txt",
			expectedFile: "./testdata/slug.out.txt",
		},
		{
			inputFile:    "./testdata/multiheadings.txt",
			expectedFile: "./testdata/multiheadings.out.txt",
		},
		{
			inputFile:    "./testdata/skipheadings.txt",
			expectedFile: "./testdata/skipheadings.out.txt",
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

		s := ltag.NewTaggingBuf("#", "#")

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
