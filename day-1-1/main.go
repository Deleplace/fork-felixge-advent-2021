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
	var prev int16 = -9999
	var increases int = -1

	// Read 1st char of 1st number
	c := input[0]
	val := int16(c - '0')
mainloop:
	for p, N := 1, len(input); p < N; p++ {
		c := input[p]
		//fmt.Printf("input[%d] is %c (%d)\n", p, c, c)
		switch {
		case c == '\n':
			if val > prev {
				//fmt.Println("Found", val, ">", prev, " !!!!!")
				increases++
			}
			//fmt.Println(prev, val, increases, " p =", p)
			prev = val
			// Read 1st char of next number
			p++
			if p >= N {
				break mainloop
			}
			c := input[p]
			val = int16(c - '0')
		case c >= '0' && c <= '9':
			val = val*10 + int16(c-'0')
		default:
			panic("A")
		}
	}
	if val > prev {
		increases++
	}
	return increases, nil
}
