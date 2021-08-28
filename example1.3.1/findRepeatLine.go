package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	CountRepeatLine := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() != "exit" {
			CountRepeatLine[input.Text()]++
		} else { // } must before else at the same line
			break
		}
	}
	for line, counts := range CountRepeatLine { //:= must be together ,without " "
		if counts > 1 {
			fmt.Printf("%s repeat %d times", line, counts) //use printf not println
			fmt.Println(" ")
		}
	}
}
