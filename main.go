package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/stathat/stathatgo"
	"os"
	"time"
)

var (
	stat    = flag.String("stat", "", "statistic name")
	ezkey   = flag.String("ezkey", "", "EZ key")
	period  = flag.Int("period", 60, "seconds between stat updates")
	verbose = flag.Bool("v", false, "verbose logging")
)

func main() {
	flag.Parse()

	if *stat == "" {
		fmt.Fprintln(os.Stderr, "stat must be specified")
		os.Exit(2)
	}

	if *ezkey == "" {
		fmt.Fprintln(os.Stderr, "ezkey must be specified")
		os.Exit(2)
	}

	lineCounter := readlines()
	tick := time.Tick(time.Duration(*period) * time.Second)
	n := 0
	for {
		select {
		case <-lineCounter:
			n++
		case <-tick:
			stathat.PostEZCount(*stat, *ezkey, n)
			if *verbose {
				fmt.Printf("%s: %d lines\n", time.Now().Format("2006-01-02 15:04:05"), n)
			}
			n = 0
		}
	}
}

func readlines() chan bool {
	c := make(chan bool)
	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			_, err := r.ReadBytes('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, "ReadBytes error:", err)
				// Overkill?  Assumes process will be automatically restarted on exit.
				os.Exit(1)
			}
			c <- true
		}
	}()
	return c
}
