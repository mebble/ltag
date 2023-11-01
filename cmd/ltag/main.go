package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mebble/ltag/internal/ltag"
	"os"
)

func main() {
	ip := flag.String("ip", ltag.DefaultIPattern, "Input pattern: string used to identify \"headings\", i.e. lines that will become tags. If a line starts with multiple occurrences of this string, that line will be considered a \"sub-heading\"")
	lp := flag.String("lp", ltag.DefaultIPattern, "Inline pattern: string used to identify inline tags, i.e. strings at the end of each line that will become tags")
	op := flag.String("op", ltag.DefaultOPattern, "Output pattern: string that specifies the format of the tags. \"$\" is the tag placeholder")
	trim := flag.Bool("trim", false, "Trim off the tags from lines that have been ltagged")

	flag.Parse()

	var operation ltag.Operation
	if *trim {
		operation = ltag.NewTrimmingBuf(*ip)
	} else {
		operation = ltag.NewTaggingBuf(*ip, *lp, *op)
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
