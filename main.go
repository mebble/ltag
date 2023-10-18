package main

import (
	"bufio"
	"fmt"
	"os"
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

		// We never get the error when running run1.sh nor run2.sh
		// Through `source ./benchmark/<run>.sh`
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(fmt.Sprintf("\nn: %d\tlen(line): %d", n, len(line)))
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("\nn: %d\tlen(line): %d", n, len(line)))
	}
}
