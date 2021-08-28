package main

import (
	"fmt"
	"os"
)

func main() {
	var cmdArgs, step string
	for i := 1; i < len(os.Args); i++ { //err i from 1 ,not 0 .os.args is a string slice
		step = " "                            //it's ok to put before or after next line
		cmdArgs = cmdArgs + os.Args[i] + step //a string is a arg

	}
	fmt.Println(cmdArgs)
}
