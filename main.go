package main

import (
	"bufio"
	"fmt"
	ltag "github.com/mebble/ltag/src"
	"os"
)

func main() {
	buf := ltag.NewStreamBuffer("#", "#")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tagged, ok := buf.Transform(line)
		if ok {
			fmt.Println(tagged)
		}
	}
}
