package main

import "flag"

// var isLogin = flag.Bool("l", false, "user login status")
// var userToken = flag.String("t", "", "Token for verify")

var IsLogin bool
var Token string

func main() {

	flag.BoolVar(&IsLogin, "l", false, "user login status")
	flag.StringVar(&Token, "t", "", "Token for verify")
	flag.Parse() // cmf input -l repesent true,cannot use -l true
	// println(*isLogin, *userToken)
	println(IsLogin, Token)
	println(flag.Args())

}
