package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lineCount := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n ", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			lineCount[line]++
		}
	}
	for line, times := range lineCount {
		if times > 1 {
			fmt.Printf("%soccur%d times\n", line, times)
		}
	}
}
