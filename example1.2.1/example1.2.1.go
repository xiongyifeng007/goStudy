package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//cmdArgs, step = "", " "
	var cmdArgs, step string
	for _, value := range os.Args[1:] { //grammer :no [1:] is ok
		step = " "
		cmdArgs += value + step
	}
	fmt.Println(cmdArgs, "len", len(cmdArgs))
	map1 := map[string]string{"name": "xyf", "age": "18"}
	fmt.Println(len(map1))
	fmt.Println(strings.Join(os.Args[1:], " ")) //add " " to every  slice element

}
