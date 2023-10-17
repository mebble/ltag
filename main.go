package main

import (
	"bufio"
	"fmt"
	ltag "github.com/mebble/ltag/src"
	"os"
)

func main() {
	out := ltag.Output{Sections: []ltag.Section{}}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		out.ParseLine(line)
	}

	fmt.Println(out.Sections)
}
