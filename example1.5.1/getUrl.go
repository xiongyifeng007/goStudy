package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		responseContent, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		a := string(responseContent)
		fmt.Println(a)
		// fmt.Printf("%s", responseContent)
		//can not string(responseContent)

	}
}
