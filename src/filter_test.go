package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterTransform(t *testing.T) {
	tests := []struct {
		inputFile  string
		expectedFile string
	}{
		{
			inputFile: "../fixtures/inline.out.txt",
			expectedFile: "../fixtures/inline.final.txt",
		},
		{
			inputFile: "../fixtures/noheadings.out.txt",
			expectedFile: "../fixtures/noheadings.final.txt",
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

		s := NewFilterBuf("#")

		inputScanner := bufio.NewScanner(inputFile)
		expectedScanner := bufio.NewScanner(expectedFile)
		for inputScanner.Scan() {
			line := inputScanner.Text()
			taggedLine, ok := s.Transform(line)

			// in reality it's always ok = true for filterTransform
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
