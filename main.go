package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		hasMore := scanner.Scan()
		if !hasMore {
			break
		}
		line := scanner.Text()
		n, err := os.Stdout.Write([]byte(line))
		// n, err := fmt.Println(line)

		// We never get the error when running run1.sh nor run2.sh
		// Through `source ./benchmark/<run>.sh`
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(fmt.Sprintf("\nn: %d\tlen(line): %d", n, len(line)))
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("\nn: %d\tlen(line): %d", n, len(line)))

		// This kinda shows that I don't need to detect a closed pipe when writing to stdout. This program will stop executing when the next process finishes
		time.Sleep(time.Second)
	}
	fmt.Println("all done")
}
