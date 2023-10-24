package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	ltag "github.com/mebble/ltag/src"
)

func main() {
	ip := flag.String("ip", "#", "Input pattern: string used to identify lines that will become tags")
	op := flag.String("op", "#", "Output pattern: string that will become the starting string of each tag")
	trim := flag.Bool("trim", false, "Trim off the tags from lines that have been ltagged")

	flag.Parse()

	var operation ltag.Operation = ltag.NewTaggingBuf(*ip, *op)
	if *trim {
		operation = ltag.NewTrimmingBuf(*ip)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tagged, ok := operation.Transform(line)
		if ok {
			fmt.Println(tagged)
		}
	}
}
