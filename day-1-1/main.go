package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	answer, err := Answer(strings.TrimSpace(string(input)))
	fmt.Printf("answer: %v\n", answer)
	return err
}

func Answer(input string) (int, error) {
	var prev int64 = -9999
	var increases int = -1

	// Read 1st char of 1st number
	c := input[0]
	val := int64(c)
mainloop:
	for p, N := 1, len(input); p < N; p++ {
		c := input[p]
		if c != '\n' {
			val = (val << 10) + int64(c)
		} else {
			if val > prev {
				increases++
			}
			prev = val
			// Read 1st char of next number
			p++
			if p >= N {
				break mainloop
			}
			c = input[p]
			val = int64(c)
		}
	}
	if val > prev {
		increases++
	}
	return increases, nil
}
