package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-sigint-and-run-a-cleanup-function-i
	go func() {
		fmt.Println("New thread")

		// I don't see us receiving the SIGPIPE ever, even when piping to head
		sig := <-c
		fmt.Println(fmt.Sprintf("sig: %d", sig))
		switch sig {
		case syscall.SIGPIPE:
			fmt.Println("Bye bye")
		default:
			fmt.Println("oh yeah")
		}
		wg.Done()
	}()

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

		// time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println("all done")
}
