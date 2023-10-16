package ltag

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
    tests := []struct {
        fixture string
        expected Output
    }{
        {
            fixture: "../fixtures/headings-hash.txt",
            expected: Output{
                Sections: []Section{
                    {"animals", []string{"cats", "tigers", "dogs", "wolves"}},
                    {"things", []string{"pen", "book", "table"}},
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

