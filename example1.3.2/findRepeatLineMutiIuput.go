package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	LineCounts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		count(os.Stdin, LineCounts)
	} else {
		for _, fileName := range file {
			fi, err := os.Open(fileName)
			if err != nil {
				fmt.Println(err)
			} else {
				count(fi, LineCounts)
				fi.Close() // remember close
			}
		}
	}
	for line, times := range LineCounts {
		if times > 1 {
			fmt.Printf("%s repeat %d times", line, times)
			fmt.Println(" ")
		}
	}
}

func count(f *os.File, LineCounts map[string]int) { //first parameter should be os.File
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() != "exit" {
			LineCounts[input.Text()]++
		} else {
			break
		}
		// for line, times := range LineCounts {          can not occur at here
		// 	if times > 1 {
		// 		fmt.Printf("%s repeat %d times", line, times)
		// 	}
		// }
	}
}
