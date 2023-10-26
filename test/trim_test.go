package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/mebble/ltag/internal/ltag"
	"github.com/stretchr/testify/assert"
)

func TestTrimTransform(t *testing.T) {
	tests := []struct {
		inputFile    string
		expectedFile string
	}{
		{
			inputFile:    "./testdata/inline.out.txt",
			expectedFile: "./testdata/inline.final.txt",
		},
		{
			inputFile:    "./testdata/noheadings.out.txt",
			expectedFile: "./testdata/noheadings.final.txt",
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

		s := ltag.NewTrimmingBuf("#")

		inputScanner := bufio.NewScanner(inputFile)
		expectedScanner := bufio.NewScanner(expectedFile)
		for inputScanner.Scan() {
			line := inputScanner.Text()
			taggedLine, ok := s.Transform(line)

			// in reality it's always ok = true for trimmingTransform
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
