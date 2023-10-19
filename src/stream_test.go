package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
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

		s := StreamBuffer{}

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
